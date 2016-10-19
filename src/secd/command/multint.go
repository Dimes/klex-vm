package command

import (
	"secd"
)

type multInt struct {
}

func MultInt() *multInt {
	return &multInt {}
}

func (mult multInt) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop().(int)
	op2 := state.S.Pop().(int)
	state.S.Push(op1 * op2)
	return state
}
