// Package node implements the Node interface and functions.
package node

import (
	"github.com/stvmln86/galah/galah/elems/icon"
)

// Node is an interface for gameworld objects.
type Node interface {
	// Icon returns the Node's Icon.
	Icon() icon.Icon

	// Name returns the Node's name.
	Name() string

	// Open returns true if the Node can be moved through.
	Open() bool
}
