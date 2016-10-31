package command

import (
	"util"
)

type Function struct {
	ArgCount int64
	Commands []Command
	Environment *util.Stack
}
