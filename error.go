package sterr

import (
	"fmt"
)

const _skip = 1

type StackTraceError struct {
	wrappedErr error
	stacktrace string
	message    string
}

func New(message string, args ...any) error {
	return &StackTraceError{
		stacktrace: takeStacktrace(_skip),
		message:    fmt.Sprintf(message, args...),
	}
}

func Wrap(wrappedErr error) error {
	if wrappedErr == nil {
		return nil
	}

	return &StackTraceError{
		stacktrace: takeStacktrace(_skip),
		wrappedErr: wrappedErr,
	}
}

func Wrapf(wrappedErr error, message string, args ...any) error {
	if wrappedErr == nil {
		return nil
	}

	return &StackTraceError{
		stacktrace: takeStacktrace(_skip),
		wrappedErr: wrappedErr,
		message:    fmt.Sprintf(message, args...),
	}
}

func (e StackTraceError) Error() string {
	switch {
	case e.wrappedErr == nil:
		return fmt.Sprintf("%s [%s]", e.stacktrace, e.message)
	case len(e.message) == 0:
		var we = e.wrappedErr.Error()
		return fmt.Sprintf("%s -> %s", e.stacktrace, we)
	default:
		var we = e.wrappedErr.Error()
		return fmt.Sprintf("%s [%s] -> %s", e.stacktrace, e.message, we)
	}
}
