// Package base implements the BaseNode type and methods.
package base

import (
	"github.com/stvmln86/galah/galah/terms/cell"
)

// BaseNode is a basic Node implementation.
type BaseNode struct {
	cell *cell.Cell
	name string
	open bool
}

// New returns a new BaseNode.
func New(cell *cell.Cell, name string, open bool) *BaseNode {
	return &BaseNode{cell, name, open}
}

// Cell returns the BaseNode's Cell.
func (n *BaseNode) Cell() *cell.Cell {
	return n.cell
}

// Name returns the BaseNode's name.
func (n *BaseNode) Name() string {
	return n.name
}

// Open returns true if the BaseNode can be moved through.
func (n *BaseNode) Open() bool {
	return n.open
}
