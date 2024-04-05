package openai

import (
	"errors"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"

	"private-llm-backend/pkg/errorutil"
)

func NewClientWithApiKey(apiKey string) (ClientWithResponsesInterface, error) {
	authProvider, err := securityprovider.NewSecurityProviderBearerToken(apiKey)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to create a new security provider"))
	}
	client, err := NewClientWithResponses(
		Host,
		WithRequestEditorFn(authProvider.Intercept),
	)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to create a new client"))
	}
	return client, nil
}
