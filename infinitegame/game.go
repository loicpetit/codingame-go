package infinitegame

import (
	"github.com/loicpetit/codingame-go/simplegame"
)

type Game struct {
	simplegame.Game
}

func (Game) Winner(state *simplegame.State) int {
	return 0
}
