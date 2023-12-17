package simplegame

import (
	"fmt"
)

type Writer struct{}

func (Writer) Write(action *Action) {
	if action == nil {
		panic("Action required to write")
	}
	fmt.Println(action.count)
}
