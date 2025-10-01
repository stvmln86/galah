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

// Add returns a new Pair with an added Pair.
func (p *Pair) Add(pair *Pair) *Pair {
	return &Pair{p.X + pair.X, p.Y + pair.Y}
}

// AddXY returns a new Pair with added co-ordinates.
func (p *Pair) AddXY(x, y int) *Pair {
	return &Pair{p.X + x, p.Y + y}
}
