package command

import (
	"secd"
	"secd/types"
)

type mod struct {
	t types.Numeric
}

func Mod(t types.Numeric) *mod {
	return &mod {t}
}

func (m mod) Evaluate(state *secd.Secd) *secd.Secd {
	op1 := state.S.Pop()
	op2 := state.S.Pop()

	switch m.t {
	case types.Int:
		state.S.Push(op1.(int64) % op2.(int64))
	case types.Int8:
		state.S.Push(op1.(int8) % op2.(int8))
	case types.Int16:
		state.S.Push(op1.(int16) % op2.(int16))
	case types.Int32:
		state.S.Push(op1.(int32) % op2.(int32))
	}

	return state
}
