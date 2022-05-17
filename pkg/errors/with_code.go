package errors

import (
	"fmt"
	"io"
)

type withCode struct {
	cause error
	code  int
	*stack
}

func WithCode(code int, errInfo any) error {
	switch err := errInfo.(type) {
	case nil:
		return nil
	case *withCode:
		return &withCode{
			code:  code,
			cause: err.cause,
			stack: err.stack,
		}
	case *withStack:
		return &withCode{
			code:  code,
			cause: err.error,
			stack: err.stack,
		}
	case error:
		return &withCode{
			code:  code,
			cause: err,
			stack: callers(),
		}
	default:
		return &withCode{
			code:  code,
			cause: fmt.Errorf("%+v", err),
			stack: callers(),
		}
	}
}

func (w *withCode) Error() string { return w.cause.Error() }
func (w *withCode) Cause() error  { return w.cause }
func (w *withCode) Unwrap() error { return w.cause }
func (w *withCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}
