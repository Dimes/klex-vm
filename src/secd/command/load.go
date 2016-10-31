package command

import (
	"secd"
)

type load struct {
	envIndex int64
	index int64
}

func Load(envIndex int64, index int64) *load {
	return &load {envIndex, index}
}

func (ld load) Evaluate(state *secd.Secd) *secd.Secd {
	state.S.Push(state.E.At(ld.envIndex).(*secd.Environment).Get(ld.index))
	return state
}
