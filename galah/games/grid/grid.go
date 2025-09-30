// Package grid implements the Grid type and methods.
package grid

import (
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Grid is a two-dimensional array of Tiles representing a gameworld.
type Grid struct {
	Tiles [][]*tile.Tile
}

// New returns a new Grid.
func New(size int) *Grid {
	var tiles [][]*tile.Tile
	for range size {
		tiles = append(tiles, make([]*tile.Tile, size))
	}

	return &Grid{tiles}
}

// Get returns a Tile from a location in the Grid.
func (g *Grid) Get(orig *pair.Pair) *tile.Tile {
	return g.Tiles[orig.Y][orig.X]
}

// Render returns the Grid as a two-dimensional Flag buffer.
func (g *Grid) Render() [][]flag.Flag {
	var flags [][]flag.Flag

	for _, line := range g.Tiles {
		flags = append(flags, make([]flag.Flag, 0))

		for _, tile := range line {
			if tile == nil {
				flags[len(flags)-1] = append(flags[len(flags)-1], " DD")
			} else {
				flags[len(flags)-1] = append(flags[len(flags)-1], tile.Flag())
			}
		}
	}

	return flags
}

// Set sets a Tile into a location in the Grid.
func (g *Grid) Set(dest *pair.Pair, tile *tile.Tile) {
	g.Tiles[dest.Y][dest.X] = tile
}
