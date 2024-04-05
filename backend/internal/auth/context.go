package auth

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"

	"private-llm-backend/internal/api"
	"private-llm-backend/pkg/errorutil"
)

var (
	ErrJWTClaimsNotFound = errors.New("jwt claims not found")
)

const jwtClaimsContextKey = "auth/jwt_claims"

func GetJWTClaimsWithEcho(c echo.Context) (*JWTClaims, error) {
	claims, ok := c.Get(jwtClaimsContextKey).(*JWTClaims)
	if !ok {
		return nil, errorutil.Error(ErrJWTClaimsNotFound)
	}
	return claims, nil
}

// GetJWTClaims retrieves the JWTClaims from within ctx of request
func GetJWTClaims(ctx context.Context) (*JWTClaims, error) {
	eCtx, err := api.GetEchoContext(ctx)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get echo context"))
	}
	claims, err := GetJWTClaimsWithEcho(eCtx)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to get jwt claims from context"))
	}
	return claims, nil
}
