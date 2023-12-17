package runner

import (
	"fmt"
	"os"
)

type OutputWriter[ACTION any] interface {
	Write(action *ACTION)
}

func WriteOutput(a ...any) {
	fmt.Println(a...)
}

func WriteDebug(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}
