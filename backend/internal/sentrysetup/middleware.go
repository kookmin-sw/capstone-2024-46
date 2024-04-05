package sentrysetup

import (
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"

	"private-llm-backend/internal/auth"
)

func SetSentryExtraContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtClaims, err := auth.GetJWTClaimsWithEcho(c)
		if err == nil {
			if hub := sentryecho.GetHubFromContext(c); hub != nil {
				hub.Scope().SetUser(sentry.User{
					ID:    jwtClaims.UserID,
					Email: jwtClaims.Email,
				})
			}
		}
		return next(c)
	}
}
