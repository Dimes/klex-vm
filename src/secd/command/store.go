package command

import (
	"secd"
)

type store struct {
	envIndex int
	index int
}

func Store(envIndex int, index int) *store {
	return &store {envIndex, index}
}

func (s store) Evaluate(state *secd.Secd) *secd.Secd {
	val := state.S.Pop()
	state.E.At(s.envIndex).(*secd.Environment).Set(s.index, val)
	return state
}
