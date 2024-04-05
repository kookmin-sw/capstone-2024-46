package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Config struct {
	ENV                  string  `json:"ENV" mapstructure:"ENV"`
	AWSSecretManagerName *string `json:"AWS_SECRET_MANAGER_NAME" mapstructure:"AWS_SECRET_MANAGER_NAME"`
	AWSRegion            *string `json:"AWS_REGION" mapstructure:"AWS_REGION"`
	SentryDsn            *string `json:"SENTRY_DSN" mapstructure:"SENTRY_DSN"`
	MySqlDsn             *string `json:"MYSQL_DSN" mapstructure:"MYSQL_DSN"`
	Port                 *string `json:"PORT" mapstructure:"PORT"`
	JWTSecret            *string `json:"JWT_SECRET" mapstructure:"JWT_SECRET"`
	ClientKey            *string `json:"CLIENT_KEY" mapstructure:"CLIENT_KEY"`
	OpenAIAPIKey         *string `json:"OPENAI_API_KEY" mapstructure:"OPENAI_API_KEY"`
	AssistantID          *string `json:"ASSISTANT_ID" mapstructure:"ASSISTANT_ID"`

	isSecretLoaded bool
}

// AllowOverwriteEnvKeys returns a list of environment keys that can be overwritten in the Config struct.
func (c *Config) AllowOverwriteEnvKeys() []string {
	val := reflect.TypeOf(Config{})
	keys := make([]string, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		tag := val.Field(i).Tag.Get("mapstructure")
		parts := strings.Split(tag, ",")
		keys = append(keys, parts[0])
	}
	return keys
}

func (c *Config) IsNeedToLoadSecret() bool {
	return !c.isSecretLoaded &&
		isNotEmpty(c.AWSRegion) && isNotEmpty(c.AWSSecretManagerName) &&
		(isEmpty(c.MySqlDsn) || isEmpty(c.JWTSecret) || isEmpty(c.ClientKey))
}

func (c *Config) Validate() error {
	if isEmpty(c.AWSRegion) {
		return errors.New("AWS_REGION is required")
	}
	if c.IsNeedToLoadSecret() {
		if isEmpty(c.AWSSecretManagerName) {
			return errors.New("AWS_SECRET_MANAGER_NAME is required")
		}
		return nil
	}
	if isEmpty(c.MySqlDsn) {
		return errors.New("MYSQL_DSN is required")
	}
	if isEmpty(c.JWTSecret) {
		return errors.New("JWT_SECRET is required")
	}
	if isEmpty(c.ClientKey) {
		return errors.New("CLIENT_KEY is required")
	}
	if isEmpty(c.OpenAIAPIKey) {
		return errors.New("OPENAI_API_KEY is required")
	}
	if isEmpty(c.AssistantID) {
		return errors.New("ASSISTANT_ID is required")
	}
	return nil
}

func (c *Config) ListenAddress(defaultPort int) string {
	if c == nil || c.Port == nil || *c.Port == "" {
		return fmt.Sprintf(":%d", defaultPort)
	}
	return fmt.Sprintf(":%s", *c.Port)
}

func isNotEmpty(s *string) bool {
	return s != nil && *s != ""
}

func isEmpty(s *string) bool {
	return !isNotEmpty(s)
}
