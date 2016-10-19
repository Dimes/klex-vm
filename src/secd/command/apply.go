package command

import (
	"util"
	"secd"
)

type apply struct {
}

func Apply() *apply {
	return &apply {}
}

func (ap apply) Evaluate(state *secd.Secd) *secd.Secd {
	dump := secd.NewDump(state.S, state.E, state.C)
	state.D.Push(dump)

	function := state.S.Pop().(Function)
	state.E = function.Environment

	env := secd.NewEnvironment()
	state.E.Push(env)
	for i := 0; i < function.ArgCount; i++ {
		env.Add(state.S.Pop())
	}

	state.C = util.NewStack()
	for i := len(function.Commands) - 1; i >= 0; i-- {
		state.C.Push(function.Commands[i])
	}

	state.S = util.NewStack()
	return state
}
