package command

import (
	"secd"
	"secd/types"
)

type add struct {
	t types.Numeric
}

func Add(t types.Numeric) *add {
	return &add {t}
}

func (a add) Evaluate(state *secd.Secd) *secd.Secd {	
	switch a.t {
	case types.Int:
		op1 := state.S.Pop().(int64)
		op2 := state.S.Pop().(int64)
		state.S.Push(op1 + op2)
	case types.Int8:
		op1 := state.S.Pop().(int8)
		op2 := state.S.Pop().(int8)
		state.S.Push(op1 + op2)
	case types.Int16:
		op1 := state.S.Pop().(int16)
		op2 := state.S.Pop().(int16)
		state.S.Push(op1 + op2)
	case types.Int32:
		op1 := state.S.Pop().(int32)
		op2 := state.S.Pop().(int32)
		state.S.Push(op1 + op2)
	case types.Float:
		op1 := state.S.Pop().(float64)
		op2 := state.S.Pop().(float64)
		state.S.Push(op1 + op2)
	}
	
	return state
}
