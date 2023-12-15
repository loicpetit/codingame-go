package runner

import (
	"time"
)

type Strategy[STATE any, ACTION any] interface {
	findAction(state *STATE, player int, maxTime time.Time) *ACTION
}
