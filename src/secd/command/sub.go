package command

import (
	"secd"
	"secd/types"
)

type sub struct {
	t types.Numeric
}

func Sub(t types.Numeric) *sub {
	return &sub {t}
}

func (s sub) Evaluate(state *secd.Secd) *secd.Secd {
	switch s.t {
	case types.Int:
		op1 := state.S.Pop().(int64)
		op2 := state.S.Pop().(int64)
		state.S.Push(op1 - op2)
	case types.Int8:
		op1 := state.S.Pop().(int8)
		op2 := state.S.Pop().(int8)
		state.S.Push(op1 - op2)
	case 16:
		op1 := state.S.Pop().(int16)
		op2 := state.S.Pop().(int16)
		state.S.Push(op1 - op2)
	case 32:
		op1 := state.S.Pop().(int32)
		op2 := state.S.Pop().(int32)
		state.S.Push(op1 - op2)
	case types.Float:
		op1 := state.S.Pop().(float64)
		op2 := state.S.Pop().(float64)
		state.S.Push(op1 - op2)
	}
	
	return state
}
