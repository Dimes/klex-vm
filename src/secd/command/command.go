package command

import (
	"secd"
)
	
type Command interface {
	Evaluate(state *secd.Secd) *secd.Secd
}
