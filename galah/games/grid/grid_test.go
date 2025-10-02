package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	xCell = cell.New('A', 'B')
	xTile = tile.New(xCell, nil)
)

func TestNew(t *testing.T) {
	// success
	grid := New(1, xTile)
	assert.Equal(t, 1, grid.Size)
	assert.Equal(t, []*tile.Tile{xTile}, grid.Tiles)
}

func TestCells(t *testing.T) {
	// setup
	grid := New(1, xTile)

	// success
	cells := grid.Cells()
	assert.Equal(t, []*cell.Cell{xCell}, cells)
}

func TestGetTile(t *testing.T) {
	// setup
	grid := New(3, nil)
	grid.Tiles[4] = xTile

	// success - in bounds
	tile := grid.GetTile(1, 1)
	assert.Equal(t, xTile, tile)

	// success - out of bounds
	tile = grid.GetTile(-1, -1)
	assert.Nil(t, tile)
}

func TestSetTile(t *testing.T) {
	// setup
	grid := New(3, nil)

	// success - in bounds
	okay := grid.SetTile(1, 1, xTile)
	assert.Equal(t, xTile, grid.Tiles[4])
	assert.True(t, okay)

	// success - out of bounds
	okay = grid.SetTile(-1, -1, xTile)
	assert.False(t, okay)
}
