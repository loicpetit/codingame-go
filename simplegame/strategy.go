package simplegame

import (
	"time"
)

type Strategy struct{}

func (Strategy) FindAction(state *State, player int, maxTime time.Time) *Action {
	return &Action{
		Player: 1,
		Count:  2,
	}
}
