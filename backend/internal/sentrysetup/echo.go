package sentrysetup

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"

	"private-llm-backend/internal/config"
	"private-llm-backend/pkg/errorutil"
)

func UseSentry(e *echo.Echo, c *config.Config) error {
	sentryDSN := ""
	if c.SentryDsn != nil {
		sentryDSN = *c.SentryDsn
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryDSN,
		Environment:      c.ENV,
		EnableTracing:    false,
		AttachStacktrace: true,
	}); err != nil {
		return errorutil.WithDetail(err, errors.New("failed to initialize sentry"))
	}
	e.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         10 * time.Second,
	}))
	return nil
}
