package runner

type Game[STATE any, ACTION any] interface {
	Start() *STATE
	GetAvailableActions(state *STATE, player int) []*ACTION
	GetNextPlayer(state *STATE) int // minimum 1
	Play(state *STATE, action *ACTION) *STATE
	Winner(state *STATE) int // 0 = no winner, 1,2,etc the winner
}
