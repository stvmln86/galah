// Package grid implements the Grid type and methods.
package grid

import (
	"github.com/stvmln86/galah/galah/elems/icon"
	"github.com/stvmln86/galah/galah/elems/node"
)

// Grid is a square grid of Nodes representing a gameworld.
type Grid struct {
	Size  int
	Nodes []node.Node
}

// New returns a new Grid.
func New(size int, nfun node.Function) *Grid {
	var nodes = make([]node.Node, size*size)
	for i := range size * size {
		nodes[i] = nfun()
	}

	return &Grid{size, nodes}
}

// Get returns the Node at a location in the Grid.
func (g *Grid) Get(x, y int) node.Node {
	if !g.In(x, y) {
		return nil
	}

	return g.Nodes[x*g.Size+y]
}

// Icons returns the Grid as a Node slice.
func (g *Grid) Icons() []icon.Icon {
	var icons []icon.Icon
	for _, node := range g.Nodes {
		icons = append(icons, node.Icon())
	}

	return icons
}

// In returns true if the Grid contains a location.
func (g *Grid) In(x, y int) bool {
	return x >= 0 && x < g.Size && y >= 0 && y < g.Size
}

// Set sets a Node at a location in the Grid.
func (g *Grid) Set(x, y int, node node.Node) {
	if !g.In(x, y) {
		return
	}

	g.Nodes[x*g.Size+y] = node
}
