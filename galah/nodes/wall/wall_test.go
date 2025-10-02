package wall

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	mockCell = cell.New('A', 'B')
)

func TestNew(t *testing.T) {
	// success
	node := New(mockCell, "test")
	assert.Equal(t, mockCell, node.cell)
	assert.Equal(t, "test", node.name)
}

func TestCell(t *testing.T) {
	// setup
	node := New(mockCell, "test")

	// success
	cell := node.Cell()
	assert.Equal(t, mockCell, cell)
}

func TestName(t *testing.T) {
	// setup
	node := New(mockCell, "test")

	// success
	name := node.Name()
	assert.Equal(t, "test", name)
}

func TestOpen(t *testing.T) {
	// setup
	node := New(mockCell, "test")

	// success
	open := node.Open()
	assert.False(t, open)
}
