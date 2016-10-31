package command

import (
	"secd"
)

type or struct {
}

func Or() *or {
	return &or {}
}

func (o or) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop().(bool)
	op2 := state.S.Pop().(bool)
	state.S.Push(op1 || op2)
	return state
}
