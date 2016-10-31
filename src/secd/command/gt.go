package command

import (
	"secd"
	"secd/types"
)

type gt struct {
	t types.Numeric
}

func Gt(t types.Numeric) *gt {
	return &gt {t}
}

func (g gt) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop()
	op2 := state.S.Pop()

	switch g.t {
	case types.Int:
		state.S.Push(op1.(int64) > op2.(int64))
	case types.Int8:
		state.S.Push(op1.(int8) > op2.(int8))
	case types.Int16:
		state.S.Push(op1.(int16) > op2.(int16))
	case types.Int32:
		state.S.Push(op1.(int32) > op2.(int32))
	case types.Float:
		state.S.Push(op1.(float64) > op2.(float64))
	}

	return state
}
