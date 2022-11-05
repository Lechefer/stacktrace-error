package sterr

import (
	"fmt"
)

const _skip = 1

type CustomError struct {
	wrappedErr error
	stacktrace string
	message    string
}

func New(message string, args ...any) *CustomError {
	return &CustomError{
		stacktrace: takeStacktrace(_skip),
		message:    fmt.Sprintf(message, args...),
	}
}

func Wrap(wrappedErr error) *CustomError {
	if wrappedErr == nil {
		return nil
	}

	return &CustomError{
		stacktrace: takeStacktrace(_skip),
		wrappedErr: wrappedErr,
	}
}

func Wrapf(wrappedErr error, message string, args ...any) *CustomError {
	if wrappedErr == nil {
		return nil
	}

	var err = CustomError{
		stacktrace: takeStacktrace(_skip),
		wrappedErr: wrappedErr,
		message:    fmt.Sprintf(message, args...),
	}

	return &err
}

func (e CustomError) Error() string {
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
