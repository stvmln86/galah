// Package plot implements two-dimensional plotting functions.
package plot

// Block returns a two-dimensional integer slice for a solid rectangle.
func Block(x1, y1, x2, y2 int) [][]int {
	var pairs [][]int
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			pairs = append(pairs, []int{x, y})
		}
	}

	return pairs
}

// Neighbours returns a two-dimensional integer slice for a point's neighbours.
func Neighbours(x, y int) [][]int {
	return [][]int{
		{x, y - 1}, {x + 1, y},
		{x, y + 1}, {x - 1, y},
	}
}

// Rectangle returns a two-dimensional integer slice for a hollow rectangle.
func Rectangle(x1, y1, x2, y2 int) [][]int {
	var pairs [][]int

	for x := x1; x <= x2; x++ {
		pairs = append(pairs, []int{x, y1})
		pairs = append(pairs, []int{x, y2})
	}

	for y := y1; y <= y2; y++ {
		pairs = append(pairs, []int{x1, y})
		pairs = append(pairs, []int{x2, y})
	}

	return pairs
}
