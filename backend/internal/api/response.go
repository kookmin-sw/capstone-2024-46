package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"

	"private-llm-backend/pkg/pointerutil"
)

func NewErrorBodyWithMessage(err error, detailMessage string) Status {
	s := NewErrorBody(err)
	var message string
	if s.Message != nil {
		message = *s.Message
	}
	if message != "" {
		message = fmt.Sprintf("%s\n%s", message, detailMessage)
	} else {
		message = detailMessage
	}
	s.Message = &message
	return s
}

func NewErrorBody(err error) Status {
	var apiErr standardAPIError
	if errors.As(err, &apiErr) {
		return Status(apiErr)
	}
	echoErr := &echo.HTTPError{}
	if errors.As(err, &echoErr) {
		e := echoStandardError(*echoErr)
		return Status{
			Code:    pointerutil.Int32(int32(e.StatusCode())),
			Message: pointerutil.String(e.ErrorMessage()),
		}
	}

	securityErr := &openapi3filter.SecurityRequirementsError{}
	if errors.As(err, &securityErr) {
		return Status{
			Code:    pointerutil.Int32(int32(codes.Unauthenticated)),
			Message: pointerutil.String(securityErr.Error()),
		}
	}

	return Status{
		Code:    pointerutil.Int32(int32(codes.Internal)),
		Message: pointerutil.String(err.Error()),
	}
}

func SetCookie(ctx context.Context, name string, value string, expireTime *time.Time) {
	eCtx, err := GetEchoContext(ctx)
	if err != nil {
		return
	}
	clientInfo := GetClientInfo(eCtx)

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	}

	if clientInfo.Proto() == "https" {
		cookie.Secure = true
		cookie.SameSite = http.SameSiteNoneMode
	}

	domain := clientInfo.ReqURL.Hostname()
	if domain != "localhost" && domain != "127.0.0.1" && 2 <= strings.Count(domain, ".") {
		domain = domain[strings.Index(domain, "."):]
	}
	cookie.Domain = domain

	if expireTime != nil {
		cookie.Expires = *expireTime
	}
	eCtx.SetCookie(cookie)
}

func SetRedirectHeader(ctx context.Context, redirectURL string) {
	eCtx, err := GetEchoContext(ctx)
	if err != nil {
		return
	}
	// temporary redirect
	eCtx.Response().Header().Set(echo.HeaderLocation, redirectURL)
}
