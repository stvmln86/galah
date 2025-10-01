package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var xCell = cell.New('A', 'B')

func TestNew(t *testing.T) {
	// success
	node := New(xCell, "test", false)
	assert.Equal(t, xCell, node.cell)
	assert.Equal(t, "test", node.name)
	assert.Equal(t, false, node.open)
}

func TestCell(t *testing.T) {
	// setup
	node := New(xCell, "test", false)

	// success
	cell := node.Cell()
	assert.Equal(t, xCell, cell)
}

func TestName(t *testing.T) {
	// setup
	node := New(xCell, "test", false)

	// success
	name := node.Name()
	assert.Equal(t, "test", name)
}

func TestOpen(t *testing.T) {
	// setup
	node := New(xCell, "test", false)

	// success
	open := node.Open()
	assert.Equal(t, false, open)
}
