package command

import (
	"secd"
)

type stop struct {
}

func Stop() *stop {
	return &stop {}
}

func (s stop) Evaluate(state *secd.Secd) *secd.Secd {
	state.Stop()
	return state
}
