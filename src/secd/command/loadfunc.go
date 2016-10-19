package command

import (
	"util"
	"secd"
)

type loadFunc struct {
	argCount int
	commands []Command
}

func LoadFunc(argCount int, commands []Command) *loadFunc {
	return &loadFunc {argCount, commands}
}

func (ldFunc loadFunc) Evaluate(state *secd.Secd) *secd.Secd {
	state.S.Push(Function {
		ldFunc.argCount,
		ldFunc.commands,
		util.CopyStack(state.E),
	})
	return state
}
