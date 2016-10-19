package command

import (
	"secd"
)

type subInt struct {
}

func SubInt() *subInt {
	return &subInt {}
}

func (sub subInt) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop().(int)
	op2 := state.S.Pop().(int)
	state.S.Push(op1 - op2)	
	return state
}
