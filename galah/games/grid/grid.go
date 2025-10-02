// Package grid implements the Grid type and methods.
package grid

import (
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/terms/cell"
)

// Grid is a grid of Nodes representing a gameworld.
type Grid struct {
	Size  int
	Nodes []node.Node
}

// New returns a new Grid.
func New(size int) *Grid {
	var nodes = make([]node.Node, size*size)
	return &Grid{size, nodes}
}

// Cells returns the Grid as a Cell slice.
func (g *Grid) Cells() []*cell.Cell {
	var cells = make([]*cell.Cell, g.Size*g.Size)
	for i := range cells {
		if node := g.Nodes[i]; node != nil {
			cells[i] = g.Nodes[i].Cell()
		} else {
			cells[i] = nil
		}
	}

	return cells
}

// GetNode returns a Node in the Grid.
func (g *Grid) GetNode(x, y int) node.Node {
	if x < 0 || x >= g.Size || y < 0 || y >= g.Size {
		return nil
	}

	return g.Nodes[y*g.Size+x]
}

// SetNode sets a Node in the Grid.
func (g *Grid) SetNode(x, y int, t node.Node) bool {
	if x < 0 || x >= g.Size || y < 0 || y >= g.Size {
		return false
	}

	g.Nodes[y*g.Size+x] = t
	return true
}
