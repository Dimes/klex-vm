package command

import (
	"secd"
)

type addInt struct {
}

func AddInt() *addInt {
	return &addInt {}
}

func (add addInt) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop().(int)
	op2 := state.S.Pop().(int)
	state.S.Push(op1 + op2)
	return state
}
