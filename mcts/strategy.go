package main

import (
	"time"

	"github.com/loicpetit/codingame-go/runner"
)

type MctsStrategy[STATE Hashable, ACTION comparable] struct {
	mcts *MCTS[STATE, ACTION]
}

func (strategy MctsStrategy[STATE, ACTION]) FindAction(state *STATE, player int, maxTime time.Time) *ACTION {
	if state == nil {
		panic("State cannot be nil")
	}
	strategy.mcts.Search(state, maxTime)
	return strategy.mcts.GetBestAction(state)
}

func NewMctsStrategy[STATE Hashable, ACTION comparable](game runner.Game[STATE, ACTION]) runner.Strategy[STATE, ACTION] {
	return MctsStrategy[STATE, ACTION]{
		mcts: NewMCTS[STATE, ACTION](game),
	}
}
