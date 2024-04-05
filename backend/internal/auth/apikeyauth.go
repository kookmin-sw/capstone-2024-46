package auth

import (
	"context"
	"errors"

	"github.com/getkin/kin-openapi/openapi3filter"

	"private-llm-backend/internal/api"
	"private-llm-backend/pkg/errorutil"
)

var (
	ErrInvalidAPIKey = errors.New("invalid api key")
)

var _ api.Authenticator = (*apiKeyAuthenticator)(nil)

type apiKeyAuthenticator struct {
	api.AuthenticatorBase

	apiKey string
}

func (a *apiKeyAuthenticator) Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	apiKey, err := a.GetValue(input)
	if err != nil {
		return errorutil.Error(err)
	}
	if apiKey == a.apiKey {
		// Success to authenticate
		return nil

	}
	return errorutil.Error(ErrInvalidAPIKey)
}

func NewAPIKeyAuthenticator(apiKey string) api.Authenticator {
	return &apiKeyAuthenticator{
		apiKey: apiKey,
	}
}
