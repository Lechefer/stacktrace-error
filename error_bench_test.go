package main

import "testing"

func errorStackTraceBench3() error {
	err := errorStackTraceBench2()
	return Wrap(err, "wrap msg")
}

func errorStackTraceBench2() error {
	err := errorStackTraceBench1()
	return Wrap(err)
}

func errorStackTraceBench1() error {
	return New("test msg")
}

func BenchmarkStackTraceError(b *testing.B) {
	b.ReportAllocs()

	var err error
	for i := 0; i < b.N; i++ {
		err = errorStackTraceBench3()
		err.Error()
	}
}
