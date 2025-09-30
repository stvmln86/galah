package plot

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/maths/pair"
)

func TestBlock(t *testing.T) {
	// success
	pairs := Block(pair.New(0, 0), pair.New(3, 3))
	assert.Equal(t, []*pair.Pair{
		pair.New(0, 0), pair.New(0, 1),
		pair.New(0, 2), pair.New(0, 3),
		pair.New(1, 0), pair.New(1, 1),
		pair.New(1, 2), pair.New(1, 3),
		pair.New(2, 0), pair.New(2, 1),
		pair.New(2, 2), pair.New(2, 3),
		pair.New(3, 0), pair.New(3, 1),
		pair.New(3, 2), pair.New(3, 3),
	}, pairs)
}

func TestCorners(t *testing.T) {
	// success
	pairs := Corners(pair.New(0, 0), pair.New(3, 3))
	assert.Equal(t, []*pair.Pair{
		pair.New(0, 0), pair.New(3, 0),
		pair.New(0, 3), pair.New(3, 3),
	}, pairs)
}

func TestNeighbours(t *testing.T) {
	// success
	pairs := Neighbours(pair.New(1, 1))
	assert.Equal(t, []*pair.Pair{
		pair.New(1, 0), pair.New(2, 1),
		pair.New(1, 2), pair.New(0, 1),
	}, pairs)
}

func TestRectangle(t *testing.T) {
	// success
	pairs := Rectangle(pair.New(0, 0), pair.New(3, 3))
	assert.Equal(t, []*pair.Pair{
		pair.New(0, 0), pair.New(0, 3),
		pair.New(1, 0), pair.New(1, 3),
		pair.New(2, 0), pair.New(2, 3),
		pair.New(3, 0), pair.New(3, 3),
		pair.New(0, 0), pair.New(3, 0),
		pair.New(0, 1), pair.New(3, 1),
		pair.New(0, 2), pair.New(3, 2),
		pair.New(0, 3), pair.New(3, 3),
	}, pairs)
}
