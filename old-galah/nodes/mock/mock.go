// Package mock implements the Mock Node type and methods.
package mock

import "github.com/stvmln86/galah/galah/terms/flag"

// Mock is a mock Node for unit testing.
type Mock struct {
	flag flag.Flag
	wall bool
}

// New returns a new Mock.
func New(flag flag.Flag, wall bool) *Mock {
	return &Mock{flag, wall}
}

// Flag returns the Mock's drawable Flag.
func (n *Mock) Flag() flag.Flag {
	return n.flag
}

// Wall returns true if the Mock cannot be moved through.
func (n *Mock) Wall() bool {
	return n.wall
}
