package command

import (
	"secd"
)

type load struct {
	envIndex int
	index int
}

func Load(envIndex int, index int) *load {
	return &load {envIndex, index}
}

func (ld load) Evaluate(state *secd.Secd) *secd.Secd {
	state.S.Push(state.E.At(ld.envIndex).(*secd.Environment).Get(ld.index))
	return state
}
