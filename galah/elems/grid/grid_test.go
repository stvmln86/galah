package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/elems/icon"
	"github.com/stvmln86/galah/galah/elems/node"
	"github.com/stvmln86/galah/galah/icons/baseicon"
	"github.com/stvmln86/galah/galah/nodes/basenode"
	"github.com/stvmln86/galah/galah/tools/tone"
)

var (
	mockIcon = baseicon.New('i', tone.Blue)
	mockNode = basenode.New(mockIcon, "node", true)
	mockFunc = func() node.Node {
		return basenode.New(mockIcon, "func", true)
	}
)

func TestNew(t *testing.T) {
	// success
	grid := New(2, mockFunc)
	assert.Equal(t, 2, grid.Size)
	assert.Len(t, grid.Nodes, 4)
	for _, node := range grid.Nodes {
		assert.NotNil(t, node)
	}
}

func TestGet(t *testing.T) {
	// setup
	grid := New(2, mockFunc)
	grid.Nodes[3] = mockNode

	// success - in bounds
	node := grid.Get(1, 1)
	assert.Equal(t, mockNode, node)

	// success - out of bounds
	node = grid.Get(-1, -1)
	assert.Nil(t, node)
}

func TestIcons(t *testing.T) {
	// setup
	grid := New(2, mockFunc)

	// success
	icons := grid.Icons()
	assert.Equal(t, []icon.Icon{mockIcon, mockIcon, mockIcon, mockIcon}, icons)
}

func TestIn(t *testing.T) {
	// setup
	grid := New(2, mockFunc)

	// success - true
	okay := grid.In(1, 1)
	assert.True(t, okay)

	// success - false
	okay = grid.In(-1, -1)
	assert.False(t, okay)
}

func TestSet(t *testing.T) {
	// setup
	grid := New(2, mockFunc)

	// success - in bounds
	grid.Set(1, 1, mockNode)
	assert.Equal(t, mockNode, grid.Nodes[3])

	// success - out of bounds
	grid.Set(-1, -1, mockNode)
}
