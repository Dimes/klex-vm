package command

import (
	"secd"
)

type eq struct {
}

func Eq() *eq {
	return &eq {}
}

func (e eq) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop()
	op2 := state.S.Pop()
	state.S.Push(op1 == op2)
	return state
}
