package command

import (
	"secd"
)

type and struct {
}

func And() *and {
	return &and {}
}

func (a and) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop().(bool)
	op2 := state.S.Pop().(bool)
	state.S.Push(op1 && op2)
	return state
}
