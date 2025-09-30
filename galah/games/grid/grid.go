// Package grid implements the Grid type and methods.
package grid

import (
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Grid is an array of Tiles representing a gameworld.
type Grid struct {
	Size  int
	Tiles []*tile.Tile
}

// New returns a new Grid.
func New(size int) *Grid {
	var tiles []*tile.Tile
	for range size * size {
		tiles = append(tiles, tile.New(" DD"))
	}

	return &Grid{size, tiles}
}

// Get returns a Tile from a location in the Grid.
func (g *Grid) Get(orig *pair.Pair) *tile.Tile {
	return g.Tiles[orig.Y*g.Size+orig.X]
}

// Render returns the Grid as a Flag buffer.
func (g *Grid) Render() []flag.Flag {
	var flags []flag.Flag
	for _, tile := range g.Tiles {
		flags = append(flags, tile.Flag())
	}

	return flags
}

// Set sets a Tile into a location in the Grid.
func (g *Grid) Set(dest *pair.Pair, tile *tile.Tile) {
	g.Tiles[dest.Y*g.Size+dest.X] = tile
}
