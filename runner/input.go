package runner

type InputReader[INPUT any, STATE any, ACTION any] interface {
	Read() chan INPUT
	UpdateState(state *STATE, input INPUT) *STATE
	ValidateAction(action *ACTION, input INPUT)
}
