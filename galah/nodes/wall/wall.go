// Package wall implements the WallNode type and methods.
package wall

import (
	"github.com/stvmln86/galah/galah/terms/cell"
)

// WallNode is a Node that can never be moved through.
type WallNode struct {
	cell *cell.Cell
	name string
}

// New returns a new WallNode.
func New(cell *cell.Cell, name string) *WallNode {
	return &WallNode{cell, name}
}

// Cell returns the WallNode's Cell.
func (n *WallNode) Cell() *cell.Cell {
	return n.cell
}

// Name returns the WallNode's name.
func (n *WallNode) Name() string {
	return n.name
}

// Open returns false as the WallNode cannot be moved through.
func (n *WallNode) Open() bool {
	return false
}
