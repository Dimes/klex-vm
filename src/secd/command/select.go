package command

import (
	"util"
	"secd"
)

type sel struct {
	ifTrue []Command
	ifFalse []Command
}


func Select(ifTrue []Command, ifFalse []Command) *sel {
	return &sel {ifTrue, ifFalse}
}

func (s sel) Evaluate(state *secd.Secd) *secd.Secd {
	dump := secd.NewDump(state.S, nil, state.C)
	state.D.Push(dump)

	env := secd.NewEnvironment()
	state.E.Push(env)
	
	op := state.S.Pop().(bool)
	commands := s.ifFalse
	if (op) {
		commands = s.ifTrue
	}

	state.C = util.NewStack()
	for i := len(commands) - 1; i >= 0; i-- {
		state.S.Push(commands[i])
	}

	state.S = util.NewStack()
	state.S.Push(nil)
	return state
}
