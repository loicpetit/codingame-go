package runner

import (
	"time"
)

type Strategy[STATE any, ACTION any] interface {
	FindAction(state *STATE, player int, maxTime time.Time) *ACTION
}

type SimpleStrategy[STATE any, ACTION any] struct {
	game Game[STATE, ACTION]
}

func (strategy SimpleStrategy[STATE, ACTION]) FindAction(state *STATE, player int, maxTime time.Time) *ACTION {
	if state == nil {
		panic("State cannot be nil")
	}
	availableActions := strategy.game.GetAvailableActions(state, player)
	for maxTime.After(time.Now()) {
		time.Sleep(5 * time.Millisecond)
	}
	if len(availableActions) == 0 {
		return nil
	}
	return availableActions[0]
}

func NewSimpleStrategy[STATE any, ACTION any](game Game[STATE, ACTION]) Strategy[STATE, ACTION] {
	return SimpleStrategy[STATE, ACTION]{game: game}
}
