package secd

import (
	"util"
)

type Dump struct {
	S *util.Stack
	E *util.Stack
	C *util.Stack
}

func NewDump(s *util.Stack, e *util.Stack, c *util.Stack) *Dump {
	return &Dump {s, e, c}
}
