package command

import (
	"util"
)

type Function struct {
	ArgCount int
	Commands []Command
	Environment *util.Stack
}
