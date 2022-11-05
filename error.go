package sterr

import (
	"fmt"
	"strings"
)

const _skip = 1

type CustomError struct {
	wrappedErr error
	stacktrace string
	message    string
}

func New(message string) *CustomError {
	return &CustomError{
		stacktrace: takeStacktrace(_skip),
		message:    message,
	}
}

func Wrap(message any, args ...any) *CustomError {
	var err = CustomError{
		stacktrace: takeStacktrace(_skip),
	}

	switch msg := message.(type) {
	case error:
		err.wrappedErr = msg
		if len(args) > 0 {
			err.message = fmt.Sprintf(strings.TrimRight(strings.Repeat("%v ", len(args)), " "), args...)
		}
	case string:
		err.message = fmt.Sprintf(msg, args...)
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
