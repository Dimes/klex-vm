package runner

import (
	"secd"
	"secd/command"
)

func Step(secd *secd.Secd) *secd.Secd {
	return secd.C.Pop().(command.Command).Evaluate(secd)
}
