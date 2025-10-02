// Package grid implements the Grid type and methods.
package grid

import (
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/terms/cell"
)

// Grid is a grid of Tiles representing a gameworld.
type Grid struct {
	Size  int
	Tiles []*tile.Tile
}

// New returns a new Grid.
func New(size int, dflt *tile.Tile) *Grid {
	var tiles = make([]*tile.Tile, size*size)
	for i := range size * size {
		tiles[i] = dflt
	}

	return &Grid{size, tiles}
}

// Cells returns the Grid as a Cell slice.
func (g *Grid) Cells() []*cell.Cell {
	var cells = make([]*cell.Cell, g.Size*g.Size)
	for i := range cells {
		cells[i] = g.Tiles[i].Cell()
	}

	return cells
}

// GetTile returns a Tile in the Grid.
func (g *Grid) GetTile(x, y int) *tile.Tile {
	if x < 0 || x >= g.Size || y < 0 || y >= g.Size {
		return nil
	}

	return g.Tiles[y*g.Size+x-1]
}

// SetTile sets a Tile in the Grid.
func (g *Grid) SetTile(x, y int, t *tile.Tile) bool {
	if x < 0 || x >= g.Size || y < 0 || y >= g.Size {
		return false
	}

	g.Tiles[y*g.Size+x-1] = t
	return true
}
