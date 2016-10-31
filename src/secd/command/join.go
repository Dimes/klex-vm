package command

import (
	"secd"
)

type join struct {
}

func Join() *join {
	return &join {}
}

func (j join) Evaluate(state *secd.Secd) *secd.Secd {
	dump := state.D.Pop().(*secd.Dump)
	retVal := state.S.Pop()
	state.E.Pop()

	state.S = dump.S
	state.C = dump.C
	state.S.Push(retVal)

	return state
}
