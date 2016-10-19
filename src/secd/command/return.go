package command

import (
	"secd"
)

type ret struct {
}

func Return() *ret {
	return &ret {}
}

func (r ret) Evaluate(state *secd.Secd) *secd.Secd {
	dump := state.D.Pop().(*secd.Dump)
	retVal := state.S.Pop()
	state.E.Pop() // Return function environment to its original state

	state.S = dump.S
	state.E = dump.E
	state.C = dump.C
	state.S.Push(retVal)
	return state
}
