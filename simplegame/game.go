package simplegame

// First player setting count to 20 or higher wins
type Game struct{}

func (Game) Start() *State {
	return &State{
		LastPlayer: 0,
		Count:      0,
	}
}
func (Game) GetAvailableActions(state *State, player int) []*Action {
	if player != 1 && player != 2 {
		panic("Player must be 1 or 2")
	}
	return []*Action{
		{player: player, count: 1},
		{player: player, count: 2},
		{player: player, count: 3},
	}
}
func (Game) GetNextPlayer(state *State) int {
	if state == nil || state.LastPlayer != 1 {
		return 1
	}
	return 2
}
func (Game) Play(state *State, action *Action) *State {
	if state == nil || action == nil {
		return nil
	}
	return &State{
		LastPlayer: action.player,
		Count:      state.Count + action.count,
	}
}
func (Game) Winner(state *State) int {
	if state == nil {
		return 0
	}
	if state.Count >= 20 {
		return state.LastPlayer
	}
	return 0
}
