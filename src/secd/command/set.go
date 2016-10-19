package command

import (
	"secd"
)

type set struct {
	envIndex int
	index int
}

func NewSet(envIndex int, index int) *set {
	return &set {envIndex, index}
}

func (s set) Evaluate(state *secd.Secd) *secd.Secd {
	val := state.S.Pop()
	state.E.At(s.envIndex).(*secd.Environment).Set(s.index, val)
	return state
}
