// Package node implements the Node interface and functions.
package node

import (
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Node is a single abstract entity in the gameworld.
type Node interface {
	// Flag returns the Node's drawable Flag.
	Flag() flag.Flag

	// Wall returns true if the Node cannot be moved through.
	Wall() bool
}
