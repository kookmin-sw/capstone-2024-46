package api

import (
	"errors"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

// HTTPErrorHandler is the global error handler for echo
func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	httpStats := http.StatusInternalServerError
	httpError := &echo.HTTPError{}
	if errors.As(err, &httpError) {
		httpStats = httpError.Code
	}

	securityErr := &openapi3filter.SecurityRequirementsError{}
	if errors.As(err, &securityErr) {
		httpStats = http.StatusUnauthorized
		multiErr := make([]error, 0, len(securityErr.Errors))
		for _, e := range securityErr.Errors {
			if errors.Is(e, ErrEmptySecurityValue) {
				continue
			}
			multiErr = append(multiErr, e)
		}
		if len(multiErr) == 0 {
			multiErr = append(multiErr, errors.New("authentication not found"))
		}
		securityErr.Errors = multiErr
		err = securityErr
	}

	isClientError := (400 <= httpStats && httpStats < 500) || errors.Is(err, context.Canceled)
	if !isClientError {
		c.Logger().Errorf("%+v\n", err)
		if hub := sentryecho.GetHubFromContext(c); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				hub.CaptureException(err)
			})
		}
	}

	resp := NewErrorBody(err)

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(httpStats)
	} else {
		err = c.JSON(httpStats, &resp)
	}
	if err != nil {
		c.Logger().Error(err)
	}
}
