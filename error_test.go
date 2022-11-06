package sterr

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func errorStackTraceTest5() error {
	return Wrap(errors.New("other err msg"))
}

func errorStackTraceTest4() error {
	return Wrapf(errors.New("other err msg"), "wrap msg")
}

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
		return New("test msg %v %v %v %v", "str1", 123, 34.61, false)
	}

	// WARN: changing the line numbers will break the test
	tests := []struct {
		name   string
		f      func() error
		expStr string
	}{
		{name: "sterr from anonymous func", f: func() error { return New("test msg") }, expStr: "sterr.Test_ErrorStackTrace.func2:42 [test msg]"},
		{name: "sterr from var func and New formating", f: t2f, expStr: "sterr.Test_ErrorStackTrace.func1:33 [test msg str1 123 34.61 false]"},
		{name: "sterr with msg", f: errorStackTraceTest1, expStr: "sterr.errorStackTraceTest1:28 [test msg]"},
		{name: "wrap sterr", f: errorStackTraceTest2, expStr: "sterr.errorStackTraceTest2:24 -> sterr.errorStackTraceTest1:28 [test msg]"},
		{name: "wrap sterr with msg", f: errorStackTraceTest3, expStr: "sterr.errorStackTraceTest3:19 [wrap msg] -> sterr.errorStackTraceTest2:24 -> sterr.errorStackTraceTest1:28 [test msg]"},
		{name: "other err with msg", f: errorStackTraceTest4, expStr: "sterr.errorStackTraceTest4:14 [wrap msg] -> [other err msg]"},
		{name: "other err without msg", f: errorStackTraceTest5, expStr: "sterr.errorStackTraceTest5:10 -> [other err msg]"},
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
