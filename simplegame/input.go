package simplegame

type Input struct {
	count int
}

type Reader struct{}

func (Reader) Read() chan Input {
	inputs := make(chan Input)
	go func() {
		for {
			inputs <- Input{count: 1}
		}
	}()
	return inputs
}

func (Reader) UpdateState(state *State, input Input) *State {
	return &State{
		LastPlayer: 2,
		Count:      state.Count + input.count,
	}
}

func (Reader) ValidateAction(action *Action, input Input) {
	if action == nil || action.Count < 1 || action.Count > 3 {
		panic("Invalid action")
	}
}
