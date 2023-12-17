package runner

import (
	"testing"

	"github.com/loicpetit/codingame-go/simplegame"
)

func NewReader() InputReader[simplegame.Input, simplegame.State, simplegame.Action] {
	return simplegame.Reader{}
}

func TestSetStateHeight(t *testing.T) {
	runner := NewRunner[simplegame.Input, simplegame.State, simplegame.Action](
		simplegame.Game{},
		simplegame.Reader{},
		simplegame.Strategy{},
		simplegame.Writer{},
	)
	finalState := runner.Run(nil)
	if finalState == nil {
		t.Fatal("Final state should not be nil")
	}
	if finalState.LastPlayer != 1 {
		t.Errorf("Expected last player to be 1 but was %d", finalState.LastPlayer)
	}
	if finalState.Count != 21 {
		t.Errorf("Expected count to be 21 but was %d", finalState.Count)
	}
}
