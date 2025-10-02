// Package plot implements two-dimensional plotting functions.
package plot

// Block returns a two-dimensional integer slice for a solid rectangle.
func Block(x1, y1, x2, y2 int) [][]int {
	var points [][]int

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			points = append(points, []int{x, y})
		}
	}

	return points
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
	var points [][]int

	for x := x1; x <= x2; x++ {
		points = append(points, []int{x, y1})
		points = append(points, []int{x, y2})
	}

	for y := y1; y <= y2; y++ {
		points = append(points, []int{x1, y})
		points = append(points, []int{x2, y})
	}

	return points
}
