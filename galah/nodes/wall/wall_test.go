package wall

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var xCell = cell.New('A', 'B')

func TestNew(t *testing.T) {
	// success
	node := New(xCell, "test")
	assert.Equal(t, xCell, node.cell)
	assert.Equal(t, "test", node.name)
}

func TestCell(t *testing.T) {
	// setup
	node := New(xCell, "test")

	// success
	cell := node.Cell()
	assert.Equal(t, xCell, cell)
}

func TestName(t *testing.T) {
	// setup
	node := New(xCell, "test")

	// success
	name := node.Name()
	assert.Equal(t, "test", name)
}

func TestOpen(t *testing.T) {
	// setup
	node := New(xCell, "test")

	// success
	open := node.Open()
	assert.False(t, open)
}
