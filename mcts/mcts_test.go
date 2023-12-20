package main

import (
	"testing"
	"time"

	"github.com/loicpetit/codingame-go/infinitegame"
	"github.com/loicpetit/codingame-go/simplegame"
)

func TestMctsGetBestActionWithState(t *testing.T) {
	state := &simplegame.State{}
	mcts := NewMCTS[simplegame.State, simplegame.Action](simplegame.Game{})
	mcts.Search(state, time.Now().Add(70*time.Millisecond))
	for _, child := range mcts.tree["0-0"].children {
		t.Log("MCTS", child)
	}
	action := mcts.GetBestAction(state)
	if action == nil {
		t.Fatal("Action should not be nil")
	}
	if action.Player != 1 {
		t.Errorf("Expected player 1 but was %d", action.Player)
	}
	if action.Count < 1 || action.Count > 3 {
		t.Errorf("Expected count between 1 and 3 but was %d", action.Count)
	}
}

func TestMctsGetBestPlayWithoutState(t *testing.T) {
	mcts := NewMCTS[simplegame.State, simplegame.Action](nil)
	action := mcts.GetBestAction(nil)
	if action != nil {
		t.Fatal("Action should be nil")
	}
}

func TestMctsSearchWithInfiniteGame(t *testing.T) {
	done := make(chan bool)
	go func() {
		state := &simplegame.State{}
		mcts := NewMCTS[simplegame.State, simplegame.Action](infinitegame.Game{})
		mcts.simulationTimeout = 0
		mcts.simulationDepth = 0
		mcts.Search(state, time.Now().Add(70*time.Millisecond))
		done <- true
	}()
	select {
	case <-time.After(200 * time.Millisecond):
		return
	case <-done:
		t.Error("Should not be done")
	}
}

func TestMctsSearchWithInfiniteGameWithTimeout(t *testing.T) {
	done := make(chan bool)
	go func() {
		defer func() {
			if recover() == nil {
				t.Error("Panic expected")
			}
			done <- true
		}()
		state := &simplegame.State{}
		mcts := NewMCTS[simplegame.State, simplegame.Action](infinitegame.Game{})
		mcts.simulationDepth = 0
		mcts.Search(state, time.Now().Add(70*time.Millisecond))
	}()
	select {
	case <-time.After(200 * time.Millisecond):
		t.Error("Test didn't finish in time")
	case <-done:
		return
	}
}

func TestMctsSearchWithInfiniteGameWithDepth(t *testing.T) {
	done := make(chan bool)
	go func() {
		state := &simplegame.State{}
		mcts := NewMCTS[simplegame.State, simplegame.Action](infinitegame.Game{})
		mcts.simulationTimeout = 0
		mcts.Search(state, time.Now().Add(70*time.Millisecond))
		for _, child := range mcts.tree["0-0"].children {
			t.Log("MCTS", child)
		}
		done <- true
	}()
	select {
	case <-time.After(200 * time.Millisecond):
		t.Error("Test didn't finish in time")
	case <-done:
		return
	}
}
