package command

import (
	"secd"
)

type loadConst struct {
	val interface{}
}

func LoadConst(val interface{}) *loadConst {
	return &loadConst {val}
}

func (ldConst loadConst) Evaluate(state *secd.Secd) *secd.Secd {
	state.S.Push(ldConst.val)
	return state
}
