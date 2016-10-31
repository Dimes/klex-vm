package command

import (
	"secd"
	"secd/types"
)

type div struct {
	t types.Numeric
}

func Div(t types.Numeric) *div {
	return &div {t}
}

func (d div) Evaluate(state *secd.Secd) *secd.Secd {
	switch d.t {
	case types.Int:
		op1 := state.S.Pop().(int64)
		op2 := state.S.Pop().(int64)
		state.S.Push(op1 / op2)
	case types.Int8:
		op1 := state.S.Pop().(int8)
		op2 := state.S.Pop().(int8)
		state.S.Push(op1 / op2)
	case types.Int16:
		op1 := state.S.Pop().(int16)
		op2 := state.S.Pop().(int16)
		state.S.Push(op1 / op2)
	case types.Int32:
		op1 := state.S.Pop().(int32)
		op2 := state.S.Pop().(int32)
		state.S.Push(op1 / op2)
	case types.Float:
		op1 := state.S.Pop().(float64)
		op2 := state.S.Pop().(float64)
		state.S.Push(op1 / op2)
	}

	return state
}
