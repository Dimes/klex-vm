package secd

import (
	"util"
)

type Secd struct {
	S *util.Stack // Stack
	E *util.Stack // Environment
	C *util.Stack // Command
	D *util.Stack // Dump
	stopped bool
}

func (secd *Secd) Stop() {
	secd.stopped = true
}

func (secd *Secd) IsStopped() bool {
	return secd.stopped;
}

func NewSecd() *Secd {
	return &Secd {
		util.NewStack(),
		util.NewStack(),
		util.NewStack(),
		util.NewStack(),
		false,
	}
}

