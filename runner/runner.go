package runner

import (
	"time"
)

type Runner[INPUT any, STATE any, ACTION any] struct {
	game     Game[STATE, ACTION]
	reader   InputReader[INPUT, STATE, ACTION]
	strategy Strategy[STATE, ACTION]
	writer   OutputWriter[ACTION]
}

func (runner Runner[INPUT, STATE, ACTION]) Run(quit chan bool) *STATE {
	timer := NewTimer()
	inputs := runner.reader.Read()
	state := runner.game.Start()
	WriteDebug("State:", state)
	round := 0
	for {
		select {
		case input := <-inputs:
			round++
			maxTime := startTimer(timer, round)
			// opponent turn
			state = runner.reader.UpdateState(state, input)
			if runner.game.Winner(state) > 0 {
				return state
			}
			// my turn
			nextAction := runner.strategy.FindAction(state, 1, maxTime)
			runner.reader.ValidateAction(nextAction, input)
			state = runner.game.Play(state, nextAction)
			WriteDebug("State:", state)
			runner.writer.Write(nextAction)
			endTimer(timer, round)
			WriteDebug("Timer:", timer)
			if runner.game.Winner(state) > 0 {
				return state
			}
		case <-quit:
			return state
		}
	}
}

func startTimer(timer *Timer, round int) time.Time {
	if round == 1 {
		timer.StartInit()
		return timer.initStart.Add(980 * time.Millisecond)
	} else {
		timer.StartRound()
		return timer.roundStart.Add(80 * time.Millisecond)
	}
}

func endTimer(timer *Timer, round int) {
	if round == 1 {
		timer.EndInit()
	} else {
		timer.EndRound()
	}
}

func NewRunner[INPUT any, STATE any, ACTION any](
	game Game[STATE, ACTION],
	reader InputReader[INPUT, STATE, ACTION],
	strategy Strategy[STATE, ACTION],
	writer OutputWriter[ACTION],
) Runner[INPUT, STATE, ACTION] {
	return Runner[INPUT, STATE, ACTION]{
		game,
		reader,
		strategy,
		writer,
	}
}
