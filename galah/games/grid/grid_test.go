package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/terms/flag"
)

var (
	mockTile = tile.New("tWK")
)

func TestNew(t *testing.T) {
	// success
	grid := New(3)
	assert.Equal(t, [][]*tile.Tile{
		{nil, nil, nil},
		{nil, nil, nil},
		{nil, nil, nil},
	}, grid.Tiles)
}

func TestGet(t *testing.T) {
	// setup
	grid := New(3)
	grid.Tiles[0][0] = mockTile

	// success
	tile := grid.Get(pair.New(0, 0))
	assert.Equal(t, mockTile, tile)
}

func TestRender(t *testing.T) {
	// setup
	grid := New(3)
	grid.Tiles[0][0] = mockTile

	// success
	flags := grid.Render()
	assert.Equal(t, [][]flag.Flag{
		{"tWK", " DD", " DD"},
		{" DD", " DD", " DD"},
		{" DD", " DD", " DD"},
	}, flags)
}

func TestSet(t *testing.T) {
	// setup
	grid := New(3)

	// success
	grid.Set(pair.New(0, 0), mockTile)
	assert.Equal(t, mockTile, grid.Tiles[0][0])
}
