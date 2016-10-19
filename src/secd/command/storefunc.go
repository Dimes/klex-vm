package command

import (
	"util"
	"secd"
)

type storeFunc struct {
	argCount int
	commands []Command
}

func StoreFunc(argCount int, commands []Command) *storeFunc {
	return &storeFunc {argCount, commands}
}

func (function storeFunc) Evaluate(state *secd.Secd) *secd.Secd {
	state.S.Push(Function {
		function.argCount,
		function.commands,
		util.CopyStack(state.E),
	})
	return state
}
