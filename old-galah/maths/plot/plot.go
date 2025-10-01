// Package plot implements two-dimensional plotting functions.
package plot

import (
	"github.com/stvmln86/galah/galah/maths/pair"
)

// Block returns a Pair slice for a plotted solid rectangle.
func Block(orig, dest *pair.Pair) []*pair.Pair {
	var pairs []*pair.Pair

	for x := orig.X; x <= dest.X; x++ {
		for y := orig.Y; y <= dest.Y; y++ {
			pairs = append(pairs, pair.New(x, y))
		}
	}

	return pairs
}

// Corners returns a Pair slice for the top left, top right, bottom left, and bottom
// right corners of a plotted hollow rectangle.
func Corners(orig, dest *pair.Pair) []*pair.Pair {
	return []*pair.Pair{
		pair.New(orig.X, orig.Y),
		pair.New(dest.X, orig.Y),
		pair.New(orig.X, dest.Y),
		pair.New(dest.X, dest.Y),
	}
}

// Neighbours returns a Pair slice for the north, east, south and west neighbours of a
// plotted Pair.
func Neighbours(p *pair.Pair) []*pair.Pair {
	return []*pair.Pair{
		pair.New(p.X, p.Y-1),
		pair.New(p.X+1, p.Y),
		pair.New(p.X, p.Y+1),
		pair.New(p.X-1, p.Y),
	}
}

// Rectangle returns a Pair slice for a plotted hollow rectangle.
func Rectangle(orig, dest *pair.Pair) []*pair.Pair {
	var pairs []*pair.Pair

	for x := orig.X; x <= dest.X; x++ {
		pairs = append(pairs, pair.New(x, orig.Y))
		pairs = append(pairs, pair.New(x, dest.Y))
	}

	for y := orig.Y; y <= dest.Y; y++ {
		pairs = append(pairs, pair.New(orig.X, y))
		pairs = append(pairs, pair.New(dest.X, y))
	}

	return pairs
}
