package command

import (
	"secd"
)

type not struct {
}

func Not() *not {
	return &not {}
}

func (n not) Evaluate(state *secd.Secd) *secd.Secd {
	op := state.S.Pop().(bool)
	state.S.Push(!op)
	return state
}
