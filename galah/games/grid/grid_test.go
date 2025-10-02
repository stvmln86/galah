package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/nodes/base"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	xCell = cell.New('A', 'B')
	xNode = base.New(xCell, "test", false)
)

func TestNew(t *testing.T) {
	// success
	grid := New(2)
	assert.Equal(t, 2, grid.Size)
	assert.Equal(t, []node.Node{nil, nil, nil, nil}, grid.Nodes)
}

func TestIn(t *testing.T) {
	// setup
	grid := New(2)

	// success - in bounds
	okay := grid.In(1, 1)
	assert.True(t, okay)

	// success - out of bounds
	okay = grid.In(-1, -1)
	assert.False(t, okay)
}

func TestCells(t *testing.T) {
	// setup
	grid := New(2)
	grid.SetNode(1, 1, xNode)

	// success
	cells := grid.Cells()
	assert.Equal(t, []*cell.Cell{nil, nil, nil, xCell}, cells)
}

func TestGetNode(t *testing.T) {
	// setup
	grid := New(2)
	grid.SetNode(1, 1, xNode)

	// success - in bounds
	node := grid.GetNode(1, 1)
	assert.Equal(t, xNode, node)

	// success - out of bounds
	node = grid.GetNode(-1, -1)
	assert.Nil(t, node)
}

func TestSetNode(t *testing.T) {
	// setup
	grid := New(2)

	// success - in bounds
	grid.SetNode(1, 1, xNode)
	assert.Equal(t, xNode, grid.Nodes[3])

	// success - out of bounds
	grid.SetNode(-1, -1, xNode)
}
