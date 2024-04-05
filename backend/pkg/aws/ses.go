package aws

import (
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func NewSESClient(cfg *awssdk.Config) *ses.Client {
	return ses.NewFromConfig(*cfg)
}
