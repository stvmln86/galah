package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/cell"
)

var (
	mockCell = cell.New('A', 'B')
	mockNode = New(mockCell, "test", false)
)

func TestNew(t *testing.T) {
	// success
	node := New(mockCell, "test", false)
	assert.Equal(t, mockCell, node.cell)
	assert.Equal(t, "test", node.name)
	assert.Equal(t, false, node.open)
}

func TestCell(t *testing.T) {
	// success
	cell := mockNode.Cell()
	assert.Equal(t, mockCell, cell)
}

func TestName(t *testing.T) {
	// success
	name := mockNode.Name()
	assert.Equal(t, "test", name)
}

func TestOpen(t *testing.T) {
	// success
	open := mockNode.Open()
	assert.Equal(t, false, open)
}
