package api

import (
	"context"

	"github.com/labstack/echo/v4"
)

func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		reqCtx := req.Context()
		reqCtx = context.WithValue(reqCtx, echoContextKey, c)
		req = req.WithContext(reqCtx)
		c.SetRequest(req)
		return next(c)
	}
}

func NoCacheResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Response().Committed {
			return next(c)
		}
		if c.Request().Method != echo.HEAD {
			responseHeaders := c.Response().Header()
			responseHeaders.Set("Cache-Control", "no-store")
			responseHeaders.Set("Pragma", "no-cache")
		}
		return next(c)
	}
}
