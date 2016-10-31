package command

import (
	"secd"
)

type floatCast struct {
}

func FloatCast() *floatCast {
	return &floatCast {}
}

func (f floatCast) Evaluate(state *secd.Secd) *secd.Secd {
	op := state.S.Pop()

	switch op := op.(type) {
	case int8:
		state.S.Push(float64(op))
	case int16:
		state.S.Push(float64(op))
	case int32:
		state.S.Push(float64(op))
	case int64:
		state.S.Push(float64(op))
	case float64:
		state.S.Push(op)
	}
	
	return state
}
