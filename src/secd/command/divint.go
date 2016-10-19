package command

import (
	"secd"
)

type divInt struct {
}

func DivInt() *divInt {
	return &divInt {}
}

func (div divInt) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop().(int)
	op2 := state.S.Pop().(int)
	state.S.Push(op1 / op2)
	return state
}
