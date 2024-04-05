package aws

import (
	"context"
	"errors"
	"log"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"

	"private-llm-backend/pkg/errorutil"
)

func MustLoadDefaultConfig(region string) *awssdk.Config {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	c, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(region))
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to load aws config"))
		log.Fatal(err)
	}
	return &c
}
