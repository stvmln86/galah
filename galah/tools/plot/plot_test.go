package plot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	// success
	points := Block(0, 0, 3, 3)
	assert.Equal(t, [][]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
		{1, 0}, {1, 1}, {1, 2}, {1, 3},
		{2, 0}, {2, 1}, {2, 2}, {2, 3},
		{3, 0}, {3, 1}, {3, 2}, {3, 3},
	}, points)
}

func TestNeighbours(t *testing.T) {
	// success
	points := Neighbours(1, 1)
	assert.Equal(t, [][]int{
		{1, 0}, {2, 1},
		{1, 2}, {0, 1},
	}, points)
}

func TestRectangle(t *testing.T) {
	// success
	points := Rectangle(0, 0, 3, 3)
	assert.Equal(t, [][]int{
		{0, 0}, {0, 3}, {1, 0}, {1, 3},
		{2, 0}, {2, 3}, {3, 0}, {3, 3},
		{0, 0}, {3, 0}, {0, 1}, {3, 1},
		{0, 2}, {3, 2}, {0, 3}, {3, 3},
	}, points)
}
