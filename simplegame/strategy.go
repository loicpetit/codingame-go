package simplegame

import (
	"time"
)

type Strategy struct{}

func (Strategy) FindAction(state *State, player int, maxTime time.Time) *Action {
	return &Action{
		player: 1,
		count:  2,
	}
}
