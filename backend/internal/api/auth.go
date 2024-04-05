package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"

	"private-llm-backend/pkg/errorutil"
)

var (
	ErrEmptySecurityValue = errors.New("empty security value")
)

type Authenticator interface {
	Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error
}

type requestParser openapi3filter.AuthenticationInput

func (r *requestParser) getCookie(name string) string {
	cookie, err := r.RequestValidationInput.Request.Cookie(name)
	if err != nil {
		return ""
	}
	if cookie.Value == "" {
		return ""
	}
	return cookie.Value
}

func (r *requestParser) getHeader(name string) string {
	return r.RequestValidationInput.Request.Header.Get(name)
}

func (r *requestParser) getQueryParam(name string) string {
	return r.RequestValidationInput.GetQueryParams().Get(name)
}

type AuthenticatorBase struct {
}

// GetValue retrieves the value of the authentication input based on the security scheme type and location.
// It returns the value as a string or an error if the value is empty or unsupported.
// The method uses a request parser to extract the value from the appropriate source (cookie, header, query, or token).
func (a *AuthenticatorBase) GetValue(input *openapi3filter.AuthenticationInput) (string, error) {
	parser := requestParser(*input)

	switch input.SecurityScheme.Type {
	case "apiKey":
		name := input.SecurityScheme.Name

		switch input.SecurityScheme.In {
		case "cookie":
			value := parser.getCookie(name)
			if value == "" {
				return "", errorutil.Error(ErrEmptySecurityValue)
			}
			return value, nil
		case "header":
			value := parser.getHeader(name)
			if value == "" {
				return "", errorutil.Error(ErrEmptySecurityValue)
			}
			return value, nil
		case "query":
			value := parser.getQueryParam(name)
			if value == "" {
				return "", errorutil.Error(ErrEmptySecurityValue)
			}
			return value, nil
		}
	case "http":
		authHeader := parser.getHeader("Authorization")
		if authHeader == "" {
			return "", errorutil.Error(ErrEmptySecurityValue)
		}
		switch input.SecurityScheme.Scheme {
		case "bearer":
			if len(authHeader) < 8 {
				return "", errors.New("invalid bearer token")
			}
			return authHeader[7:], nil
		}
	}

	return "", errors.New("unsupported security type")
}

func NewAuthentication(authenticators map[string]Authenticator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		auth, ok := authenticators[input.SecuritySchemeName]
		if !ok {
			return fmt.Errorf(`no authenticator for security scheme "%s"`, input.SecuritySchemeName)
		}
		return auth.Authenticate(ctx, input)
	}
}
