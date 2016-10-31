package command

import (
	"secd"
	"secd/types"
)

type resize struct {
	t types.Numeric
}

func Resize(t types.Numeric) *resize {
	return &resize {t}
}

func resize8(n interface{}) int8 {
	switch n := n.(type) {
	case int8:
		return int8(n)
	case int16:
		return int8(n)
	case int32:
		return int8(n)
	case int64:
		return int8(n)
	}

	return 0
}

func resize16(n interface{}) int16 {
	switch n := n.(type) {
	case int8:
		return int16(n)
	case int16:
		return int16(n)
	case int32:
		return int16(n)
	case int64:
		return int16(n)
	}

	return 0
}

func resize32(n interface{}) int32 {
	switch n := n.(type) {
	case int8:
		return int32(n)
	case int16:
		return int32(n)
	case int32:
		return int32(n)
	case int64:
		return int32(n)
	}

	return 0
}

func resizeInt(n interface{}) int64 {
	switch n := n.(type) {
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	}

	return 0
}

func (r resize) Evaluate(state *secd.Secd) *secd.Secd {
	op := state.S.Pop()

	switch r.t {
	case types.Int:
		state.S.Push(resizeInt(op))
	case types.Int8:
		state.S.Push(resize8(op))
	case types.Int16:
		state.S.Push(resize16(op))
	case types.Int32:
		state.S.Push(resize32(op))
	}
	
	return state
}
