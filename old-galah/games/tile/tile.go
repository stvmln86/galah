// Package tile implements the Tile type and methods.
package tile

import (
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Tile is a two-dimensional location in the gameworld.
type Tile struct {
	Base  flag.Flag
	Nodes []node.Node
}

// New returns a new Tile.
func New(flag flag.Flag, nodes ...node.Node) *Tile {
	return &Tile{flag, nodes}
}

// Empty returns true if the Tile has no Nodes.
func (t *Tile) Empty() bool {
	return len(t.Nodes) == 0
}

// Flag returns the Tile's base or topmost Flag.
func (t *Tile) Flag() flag.Flag {
	if len(t.Nodes) == 0 {
		return t.Base
	}

	return t.Nodes[len(t.Nodes)-1].Flag()
}

// Pop removes and returns the Tile's topmost Node.
func (t *Tile) Pop() node.Node {
	if len(t.Nodes) == 0 {
		return nil
	}

	node := t.Nodes[len(t.Nodes)-1]
	t.Nodes = t.Nodes[:len(t.Nodes)-1]
	return node
}

// Push adds a Node to the top of the Tile.
func (t *Tile) Push(node node.Node) {
	t.Nodes = append(t.Nodes, node)
}

// Wall returns true if the Node cannot be moved through.
func (t *Tile) Wall() bool {
	if len(t.Nodes) == 0 {
		return false
	}

	return t.Nodes[len(t.Nodes)-1].Wall()
}
