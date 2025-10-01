// Package wall implements the Wall Node type and methods.
package wall

import (
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Wall is a static Node that cannot move or be moved through.
type Wall struct {
	flag flag.Flag
}

// New returns a new Wall.
func New(flag flag.Flag) *Wall {
	return &Wall{flag}
}

// Flag returns the Wall's drawable Flag.
func (n *Wall) Flag() flag.Flag {
	return n.flag
}

// Wall returns true if the Wall cannot be moved through.
func (n *Wall) Wall() bool {
	return true
}
