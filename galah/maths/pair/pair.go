// Package pair implements the Pair type and methods.
package pair

// Pair is a two-dimensional co-ordinate pair.
type Pair struct {
	X int
	Y int
}

// New returns a new Pair.
func New(x, y int) *Pair {
	return &Pair{x, y}
}
