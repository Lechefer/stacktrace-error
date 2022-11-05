package sterr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func errorStackTraceTest3() error {
	err := errorStackTraceTest2()
	return Wrapf(err, "wrap msg")
}

func errorStackTraceTest2() error {
	err := errorStackTraceTest1()
	return Wrap(err)
}

func errorStackTraceTest1() error {
	return New("test msg")
}

func Test_ErrorStackTrace(t *testing.T) {
	t2f := func() error {
		return New("test msg")
	}

	// WARN: changing the line numbers will break the test
	tests := []struct {
		name   string
		f      func() error
		expStr string
	}{
		{name: "err from anonymous func", f: func() error { return New("test msg") }, expStr: "sterr.Test_ErrorStackTrace.func2:33 [test msg]"},
		{name: "err from var func", f: t2f, expStr: "sterr.Test_ErrorStackTrace.func1:24 [test msg]"},
		{name: "err with msg", f: errorStackTraceTest1, expStr: "sterr.errorStackTraceTest1:19 [test msg]"},
		{name: "wrap err with msg", f: errorStackTraceTest2, expStr: "sterr.errorStackTraceTest2:15 -> sterr.errorStackTraceTest1:19 [test msg]"},
		{name: "wrap with msg", f: errorStackTraceTest3, expStr: "sterr.errorStackTraceTest3:10 [wrap msg] -> sterr.errorStackTraceTest2:15 -> sterr.errorStackTraceTest1:19 [test msg]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.f()
			assert.Equal(t, tt.expStr, err.Error())
		})
	}
}

func Test_Wrap(t *testing.T) {
	t.Run("wrap nil err", func(t *testing.T) {
		assert.Equal(t, nil, Wrap(nil))
	})
}

func Test_Wrapf(t *testing.T) {
	t.Run("wrapf nil err", func(t *testing.T) {
		assert.Equal(t, nil, Wrapf(nil, "test msg"))
	})

	t.Run("args", func(t *testing.T) {
		assert.Equal(t, "str1 123 34.61 false", Wrapf(New(""), "%v %v %v %v", "str1", 123, 34.61, false).(*StackTraceError).message)
	})
}
