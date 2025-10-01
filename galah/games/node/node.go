// Package node implements the Node interface and functions.
package node

import (
	"github.com/stvmln86/galah/galah/terms/cell"
)

// Node is an interface for gameworld objects.
type Node interface {
	// Cell returns the Node's Cell.
	Cell() *cell.Cell

	// Name returns the Node's name.
	Name() string

	// Open returns true if the Node can be moved through.
	Open() bool
}
