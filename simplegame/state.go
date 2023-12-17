package simplegame

import (
	"fmt"
)

type State struct {
	LastPlayer int
	Count      int
}

func (state *State) String() string {
	return fmt.Sprintf("{lastPlayer: %d, count: %d}", state.LastPlayer, state.Count)
}
