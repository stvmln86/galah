// Package tile implements the Tile type and methods.
package tile

import (
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/terms/cell"
)

// Tile is a single inhabitable location in a Grid.
type Tile struct {
	cell *cell.Cell
	node node.Node
}

// New returns a new Tile.
func New(cell *cell.Cell, node node.Node) *Tile {
	return &Tile{cell, node}
}

// Cell returns the Tile's Cell.
func (t *Tile) Cell() *cell.Cell {
	return t.cell
}

// Node returns the Tile's Node.
func (t *Tile) Node() node.Node {
	return t.node
}

// Open returns true if the Tile can be moved through.
func (t *Tile) Open() bool {
	if t.node != nil {
		return t.node.Open()
	}

	return true
}

// SetNode sets the Tile's Node.
func (t *Tile) SetNode(node node.Node) {
	t.node = node
}
