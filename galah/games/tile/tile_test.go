package tile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/nodes/base"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	mockCell = cell.New('A', 'B')
	mockNode = base.New(mockCell, "test", false)
	mockTile = func() *Tile { return New(mockCell, mockNode) }
)

func TestNew(t *testing.T) {
	// success
	tile := New(mockCell, mockNode)
	assert.Equal(t, mockCell, tile.cell)
	assert.Equal(t, mockNode, tile.node)
}

func TestCell(t *testing.T) {
	// success
	cell := mockTile().Cell()
	assert.Equal(t, mockCell, cell)
}

func TestNode(t *testing.T) {
	// success
	node := mockTile().Node()
	assert.Equal(t, mockNode, node)
}

func TestOpen(t *testing.T) {
	// setup
	tile := mockTile()

	// success - with node
	open := tile.Open()
	assert.False(t, open)

	// setup
	tile.node = nil

	// success - without node
	open = tile.Open()
	assert.True(t, open)
}

func TestSetNode(t *testing.T) {
	// setup
	tile := mockTile()
	tile.node = nil

	// success
	tile.SetNode(mockNode)
	assert.Equal(t, mockNode, tile.node)
}
