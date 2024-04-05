package auth

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3filter"

	"private-llm-backend/internal/api"
)

var _ api.Authenticator = (*cookieJWTProvider)(nil)

type cookieJWTProvider struct {
	auth api.Authenticator
}

func (a *cookieJWTProvider) Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	eCtx, err := api.GetEchoContext(ctx)
	if err != nil {
		return a.auth.Authenticate(ctx, input)
	}
	authorizationHeader := eCtx.Request().Header.Get("Authorization")
	if authorizationHeader != "" {
		// Use authorization header auth first
		return api.ErrEmptySecurityValue
	}
	return a.auth.Authenticate(ctx, input)
}

// NewCookieJWTAuthenticator 는 'Authorization' 헤더가 비어있을 때 쿠키를 사용하여 JWT 를 인증하는 Authenticator 를 생성합니다.
// 만약 'Authorization' 헤더가 비어있지 않다면, 다른 Authenticator 를 사용할 수 있도록 api.ErrEmptySecurityValue 에러를 반환합니다.
func NewCookieJWTAuthenticator(jwtProvider JWTProvider) api.Authenticator {
	auth := NewJWTAuthenticator(jwtProvider)
	return &cookieJWTProvider{
		auth: auth,
	}
}
