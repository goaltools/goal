// Package command is used for parsing input parameters.
package command

import (
	"github.com/anonx/ok/log"
)

// Type is a main type of command package.
// It is used for storage of parsed parameters.
type Type struct {
	action string
}

// NewType initializes and returns Type object.
// It expects even number of args. Otherwise, it will panic.
func NewType(args []string) Type {
	// Make sure the number of arguments is even number.
	if len(args)%2 != 0 {
		log.Error.Panic("Incorrect number of parameters.")
	}
	return Type{}
}
