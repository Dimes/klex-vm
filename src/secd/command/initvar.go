package command

import (
	"secd"
)

type initVar struct {
}

func InitVar() *initVar {
	return &initVar {}
}

func (init initVar) Evaluate(state *secd.Secd) *secd.Secd {
	env := state.E.Peek().(*secd.Environment)
	env.Add(state.S.Pop())
	return state
}
