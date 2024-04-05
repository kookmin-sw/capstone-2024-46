// An error must serve two purposes: error comparison for handling, and ease of debugging for unexpected error resolve
// This util provides a simple way to serve those.
//
// This util defines some behavior interface for behavior based handling like Causer, Retryable.
// This util automatically inject the stack trace (a slice of PC).

package errorutil

import (
	"errors"
	"fmt"
	"io"
)

type fundamental struct {
	err   error
	cause error
	*stack
}

var _ error = (*fundamental)(nil)
var _ fmt.Formatter = (*fundamental)(nil)
var _ Causer = (*fundamental)(nil)

func (f *fundamental) Error() string {
	return f.err.Error()
}

func (f *fundamental) Cause() error {
	return f.cause
}

func (f *fundamental) Is(err error) bool {
	return errors.Is(f.err, err)
}

func (f *fundamental) As(target interface{}) bool {
	return errors.As(f.err, target) || errors.As(f.cause, target)
}

// Unwrap stands for error chaining compatibility
func (f *fundamental) Unwrap() error {
	return f.cause
}

func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if f.Cause() != nil && s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%s (caused by: %v)", f.Error(), f.Cause())
			return
		}
		fallthrough
	case 's', 'q':
		_, _ = io.WriteString(s, f.Error())
	}
}

func WithDetail(cause error, err error) error {
	if err == nil {
		return nil
	}
	if fe, ok := err.(*fundamental); ok && cause == nil {
		// If error is for bypass and error is already fundamental, just return it
		fe.stack = callers()
		return fe
	}

	return &fundamental{
		err:   err,
		cause: cause,
		stack: callers(),
	}
}

func Error(err error) error {
	return WithDetail(nil, err)
}
