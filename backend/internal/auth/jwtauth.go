package auth

import (
	"context"
	"errors"

	"github.com/getkin/kin-openapi/openapi3filter"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"

	"private-llm-backend/internal/api"
	"private-llm-backend/pkg/errorutil"
)

var (
	ErrInvalidJWT = errors.New("invalid jwt token")
)

var _ api.Authenticator = (*JWTAuthenticator)(nil)

type JWTAuthenticator struct {
	api.AuthenticatorBase
	jwt JWTProvider
}

func (a *JWTAuthenticator) Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	tokenString, err := a.GetValue(input)
	if err != nil {
		return errorutil.Error(err)
	}
	claims, err := a.jwt.Verify(tokenString)
	if err != nil {
		return errorutil.Error(err)
	}

	// Set claims to context from within requests
	eCtx := oapimiddleware.GetEchoContext(ctx)
	eCtx.Set(jwtClaimsContextKey, claims)
	return nil
}

func NewJWTAuthenticator(jwt JWTProvider) api.Authenticator {
	return &JWTAuthenticator{
		jwt: jwt,
	}
}
