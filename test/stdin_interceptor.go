package test

import (
	"fmt"
	"os"
)

type StdinInterceptor struct {
	original *os.File
	reader   *os.File
	writer   *os.File
}

func (interceptor *StdinInterceptor) Intercept() {
	interceptor.original = os.Stdin
	interceptor.reader, interceptor.writer, _ = os.Pipe()
	os.Stdin = interceptor.reader
}

func (interceptor *StdinInterceptor) Write(a ...any) {
	fmt.Fprintln(interceptor.writer, a...)
}

func (interceptor *StdinInterceptor) Close() {
	if interceptor.original != nil {
		os.Stdin = interceptor.original
		interceptor.original = nil
	}
	if interceptor.writer != nil {
		interceptor.writer.Close()
		interceptor.writer = nil
	}
	if interceptor.reader != nil {
		interceptor.reader.Close()
		interceptor.reader = nil
	}
}

func NewStdinInterceptor() *StdinInterceptor {
	return &StdinInterceptor{}
}
