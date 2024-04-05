package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"

	"private-llm-backend/pkg/pointerutil"
)

var (
	ErrPermissionDenied = echo.ErrForbidden
	ErrInvalidArgument  = echo.ErrBadRequest
)

var _ error = (*standardAPIError)(nil)

type standardAPIError Status

func (s standardAPIError) Error() string {
	if s.Message == nil {
		return "unknown error"
	}
	return *s.Message
}

func NewAPIError(code codes.Code, message string) error {
	return standardAPIError{
		Message: pointerutil.String(message),
		Code:    pointerutil.Int32(int32(code)),
	}
}

type echoStandardError echo.HTTPError

// ErrorMessage returns the string error message for the given echo error.
func (e *echoStandardError) ErrorMessage() string {
	he := (*echo.HTTPError)(e)
	if he.Internal != nil {
		return he.Internal.Error()
	}
	return fmt.Sprintf("%v", he.Message)
}

// StatusCode returns the gRPC status code for the given echo error.
func (e *echoStandardError) StatusCode() codes.Code {
	switch (*echo.HTTPError)(e) {
	case echo.ErrBadRequest:
		return codes.InvalidArgument
	case echo.ErrUnauthorized:
		return codes.Unauthenticated
	case echo.ErrForbidden:
		return codes.PermissionDenied
	case echo.ErrNotFound:
		return codes.NotFound
	case echo.ErrMethodNotAllowed, echo.ErrNotAcceptable,
		echo.ErrUnsupportedMediaType, echo.ErrUpgradeRequired,
		echo.ErrPreconditionRequired, echo.ErrTooEarly,
		echo.ErrRequestHeaderFieldsTooLarge,
		echo.ErrHTTPVersionNotSupported, echo.ErrVariantAlsoNegotiates,
		echo.ErrNotExtended:
		return codes.Unimplemented
	case echo.ErrRequestTimeout:
		return codes.DeadlineExceeded
	case echo.ErrConflict:
		return codes.Aborted
	case echo.ErrGone, echo.ErrServiceUnavailable, echo.ErrBadGateway:
		return codes.Unavailable
	case echo.ErrLengthRequired, echo.ErrRequestURITooLong:
		return codes.InvalidArgument
	case echo.ErrPreconditionFailed, echo.ErrLocked, echo.ErrFailedDependency:
		return codes.FailedPrecondition
	case echo.ErrTooManyRequests, echo.ErrInsufficientStorage:
		return codes.ResourceExhausted
	case echo.ErrInternalServerError:
		return codes.Internal
	case echo.ErrProxyAuthRequired, echo.ErrNetworkAuthenticationRequired:
		return codes.Unauthenticated
	}
	return codes.Unknown
}
