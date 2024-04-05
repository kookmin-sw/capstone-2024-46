package api

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"

	"private-llm-backend/pkg/errorutil"
)

var (
	ErrContextNotFound = errors.New("context not found")
)

const echoContextKey = "api/echo"

func GetEchoContext(ctx context.Context) (echo.Context, error) {
	var eCtx echo.Context

	// In case ctx is reqCtx
	eCtx, _ = ctx.Value(echoContextKey).(echo.Context)
	if eCtx == nil {
		// In case when in validator
		eCtx = oapimiddleware.GetEchoContext(ctx)
		if eCtx == nil {
			return nil, errorutil.Error(ErrContextNotFound)

		}
	}
	return eCtx, nil
}
