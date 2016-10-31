package command

import (
	"secd"
)

type store struct {
	envIndex int64
	index int64
}

func Store(envIndex int64, index int64) *store {
	return &store {envIndex, index}
}

func (s store) Evaluate(state *secd.Secd) *secd.Secd {
	val := state.S.Pop()
	state.E.At(s.envIndex).(*secd.Environment).Set(s.index, val)
	return state
}
