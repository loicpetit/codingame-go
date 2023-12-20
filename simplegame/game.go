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
	if player == 1 {
		return []*Action{
			{Player: player, Count: 1},
			{Player: player, Count: 2},
			{Player: player, Count: 3},
		}
	}
	if player == 2 {
		return []*Action{
			{Player: player, Count: 1},
		}
	}
	panic("Player must be 1 or 2")
}
func (Game) GetLastPlayer(state *State) int {
	if state == nil {
		return 0
	}
	return state.LastPlayer
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
	if action.Player != 1 && action.Player != 2 {
		panic("Player must be 1 or 2")
	}
	if action.Count < 1 || action.Count > 3 {
		panic("Count must be between 1 and 3")
	}
	return &State{
		LastPlayer: action.Player,
		Count:      state.Count + action.Count,
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
