package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/nodes/base"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	mockNodeCell = cell.New('A', 'B')
	mockBaseCell = cell.New(' ', 'D')
	mockNode     = base.New(mockNodeCell, "test", false)
	mockFunc     = func() node.Node {
		return base.New(mockBaseCell, "", true)
	}
)

func TestNew(t *testing.T) {
	// success
	grid := New(2, mockFunc)
	assert.Equal(t, 2, grid.Size)
	assert.Len(t, grid.Nodes, 4)
}

func TestIn(t *testing.T) {
	// setup
	grid := New(2, mockFunc)

	// success - in bounds
	okay := grid.In(1, 1)
	assert.True(t, okay)

	// success - out of bounds
	okay = grid.In(-1, -1)
	assert.False(t, okay)
}

func TestCells(t *testing.T) {
	// setup
	grid := New(2, mockFunc)
	grid.SetNode(1, 1, mockNode)

	// success
	cells := grid.Cells()
	assert.Equal(t, []*cell.Cell{mockBaseCell, mockBaseCell, mockBaseCell, mockNodeCell}, cells)
}

func TestGetNode(t *testing.T) {
	// setup
	grid := New(2, mockFunc)
	grid.SetNode(1, 1, mockNode)

	// success - in bounds
	node := grid.GetNode(1, 1)
	assert.Equal(t, mockNode, node)

	// success - out of bounds
	node = grid.GetNode(-1, -1)
	assert.Nil(t, node)
}

func TestSetNode(t *testing.T) {
	// setup
	grid := New(2, mockFunc)

	// success - in bounds
	grid.SetNode(1, 1, mockNode)
	assert.Equal(t, mockNode, grid.Nodes[3])

	// success - out of bounds
	grid.SetNode(-1, -1, nil)
}
