package runner

import (
	"time"
)

type Strategy[STATE any, ACTION any] interface {
	FindAction(state *STATE, player int, maxTime time.Time) *ACTION
}
