package config

import (
	"context"
	"errors"
	"fmt"
	"os"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/spf13/viper"

	"private-llm-backend/pkg/aws"
	"private-llm-backend/pkg/errorutil"
)

func loadEnvs(keys []string) map[string]string {
	envs := make(map[string]string)
	for _, key := range keys {
		value := os.Getenv(key)
		if value != "" {
			envs[key] = value
		}
	}
	return envs
}

func deleteEnvIfExist(m map[string]string) {
	for k := range m {
		if os.Getenv(k) != "" {
			delete(m, k)
		}
	}
}

func setValues(v *viper.Viper, m map[string]string) {
	for k, val := range m {
		v.Set(k, val)
	}
}

func unmarshalConfig(v *viper.Viper, c *Config) error {
	err := v.Unmarshal(c)
	if err != nil {
		return err
	}
	return nil
}

func Load(ctx context.Context, path string, env string) (*Config, *awssdk.Config, error) {
	var v = viper.New()

	fmt.Printf("Loading config from: %s\n", env)

	v.AddConfigPath(path)
	v.SetConfigName(env)
	v.SetConfigType("json")

	c := Config{}

	// Load environment variables first
	envKeys := c.AllowOverwriteEnvKeys()
	envs := loadEnvs(envKeys)
	setValues(v, envs)

	// Load config file
	err := v.ReadInConfig()
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to read config"))
		return nil, nil, err
	}

	err = unmarshalConfig(v, &c)
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to unmarshal config"))
		return nil, nil, err
	}

	err = c.Validate()
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to validate config"))
		return nil, nil, err
	}

	awsConfig := aws.MustLoadDefaultConfig(*c.AWSRegion)

	if c.IsNeedToLoadSecret() {
		secrets, err := aws.LoadSecret(ctx, awsConfig, *c.AWSSecretManagerName)
		if err != nil {
			err = errorutil.WithDetail(err, errors.New("failed to load secret manager"))
			return nil, nil, err
		}
		deleteEnvIfExist(secrets)
		setValues(v, secrets)

		c = Config{}
		err = unmarshalConfig(v, &c)
		if err != nil {
			err = errorutil.WithDetail(err, errors.New("failed to unmarshal config after secret load"))
			return nil, nil, err
		}
		c.isSecretLoaded = true

		err = c.Validate()
		if err != nil {
			err = errorutil.WithDetail(err, errors.New("failed to validate config after secret load"))
			return nil, nil, err
		}
	}

	if c.ENV == "" {
		c.ENV = env
	}

	return &c, awsConfig, nil
}
