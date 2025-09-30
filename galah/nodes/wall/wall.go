// Package wall implements the Wall type and methods.
package wall

import (
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Wall is a solid static wall in the gameworld.
type Wall struct {
	flag flag.Flag
}

// New returns a new Wall.
func New(flag flag.Flag) *Wall {
	return &Wall{flag}
}

// Flag returns the Wall's drawable Flag.
func (w *Wall) Flag() flag.Flag {
	return w.flag
}

// Wall returns true if the Wall cannot be moved through.
func (w *Wall) Wall() bool {
	return true
}
