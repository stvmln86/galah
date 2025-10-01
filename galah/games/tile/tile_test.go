package tile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/nodes/base"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	xCell     = cell.New('A', 'B')
	xNodeCell = cell.New('N', 'R')
	xNode     = base.New(xNodeCell, "test", false)
)

func TestNew(t *testing.T) {
	// success
	tile := New(xCell, xNode)
	assert.Equal(t, xCell, tile.cell)
	assert.Equal(t, xNode, tile.node)
}

func TestCell(t *testing.T) {
	// setup
	tile := New(xCell, xNode)

	// success - with node
	cell := tile.Cell()
	assert.Equal(t, xNodeCell, cell)

	// setup
	tile = New(xCell, nil)

	// success - without node
	cell = tile.Cell()
	assert.Equal(t, xCell, cell)
}

func TestNode(t *testing.T) {
	// setup
	tile := New(xCell, xNode)

	// success
	node := tile.Node()
	assert.Equal(t, xNode, node)
}

func TestOpen(t *testing.T) {
	// setup
	tile := New(xCell, xNode)

	// success - with node
	open := tile.Open()
	assert.False(t, open)

	// setup
	tile = New(xCell, nil)

	// success - without node
	open = tile.Open()
	assert.True(t, open)
}

func TestSetNode(t *testing.T) {
	// setup
	tile := New(xCell, nil)

	// success
	tile.SetNode(xNode)
	assert.Equal(t, xNode, tile.node)
}
