// Package grid implements the Grid type and methods.
package grid

import (
	"fmt"

	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/terms/cell"
)

// Grid is a grid of Nodes representing a gameworld.
type Grid struct {
	Size  int
	Nodes []node.Node
}

// New returns a new Grid.
func New(size int, nfun func() node.Node) *Grid {
	var nodes = make([]node.Node, size*size)
	for i := range nodes {
		nodes[i] = nfun()
	}

	return &Grid{size, nodes}
}

// Cells returns the Grid as a Cell slice.
func (g *Grid) Cells() []*cell.Cell {
	var cells = make([]*cell.Cell, g.Size*g.Size)
	for i := range cells {
		fmt.Printf(">>> %d %#v\n", i, g.Nodes[i])
		cells[i] = g.Nodes[i].Cell()
	}

	return cells
}

// GetNode returns a Node in the Grid.
func (g *Grid) GetNode(x, y int) node.Node {
	if !g.In(x, y) {
		return nil
	}

	return g.Nodes[y*g.Size+x]
}

// In returns true if a co-ordinate pair is inside the Grid.
func (g *Grid) In(x, y int) bool {
	return x >= 0 && x < g.Size && y >= 0 && y < g.Size
}

// SetNode sets a Node in the Grid.
func (g *Grid) SetNode(x, y int, t node.Node) {
	if !g.In(x, y) {
		return
	}

	g.Nodes[y*g.Size+x] = t
}
