package aws

import (
	"context"
	"encoding/json"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pkg/errors"

	"private-llm-backend/pkg/errorutil"
)

func LoadSecret(ctx context.Context, cfg *awssdk.Config, secretName string) (map[string]string, error) {
	svc := secretsmanager.NewFromConfig(*cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     awssdk.String(secretName),
		VersionStage: awssdk.String("AWSCURRENT"),
	}
	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get secret value"))
	}
	m := make(map[string]string)
	err = json.Unmarshal([]byte(*result.SecretString), &m)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to unmarshal secret value"))
	}
	return m, nil
}
