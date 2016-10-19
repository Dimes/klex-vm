package command

import (
	"util"
	"secd"
)

type recApply struct {
}

func RecApply() *recApply {
	return &recApply {}
}

func (ap recApply) Evaluate(state *secd.Secd) *secd.Secd {
	dump := secd.NewDump(state.S, state.E, state.C)
	state.D.Push(dump)

	// The difference between AP and RAP is that AP replaces the current
	// environment with the function's environment. RAP, on the other hand,
	// copies the function environment. This allows functions to overwrite
	// the head of the environment with their local state without modifying
	// previous calls' environments. This operation is probably very slow.
	function := state.S.Pop().(Function)
	state.E = util.CopyStack(function.Environment)

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
