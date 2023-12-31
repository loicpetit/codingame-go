package test

import (
	"fmt"
	"os"
)

type StdoutInterceptor struct {
	original *os.File
	reader   *os.File
	writer   *os.File
}

func (interceptor *StdoutInterceptor) Intercept() {
	interceptor.original = os.Stdout
	interceptor.reader, interceptor.writer, _ = os.Pipe()
	os.Stdout = interceptor.writer
}

func (interceptor *StdoutInterceptor) Scan(a ...any) {
	fmt.Fscan(interceptor.reader, a...)
}

func (interceptor *StdoutInterceptor) Close() {
	if interceptor.original != nil {
		os.Stdout = interceptor.original
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

func NewStdoutInterceptor() *StdoutInterceptor {
	return &StdoutInterceptor{}
}
