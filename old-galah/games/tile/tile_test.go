package tile

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/nodes/mock"
	"github.com/stvmln86/galah/galah/terms/flag"
)

var (
	mockNode = mock.New("mWK", true)
)

func TestNew(t *testing.T) {
	// success
	tile := New("tWK", mockNode)
	assert.Equal(t, flag.Flag("tWK"), tile.Base)
	assert.Equal(t, []node.Node{mockNode}, tile.Nodes)
}

func TestEmpty(t *testing.T) {
	// setup
	tile := New("tWK")

	// success - true
	okay := tile.Empty()
	assert.True(t, okay)

	// setup
	tile = New("tWK", mockNode)

	// success - false
	okay = tile.Empty()
	assert.False(t, okay)
}

func TestFlag(t *testing.T) {
	// setup
	tile := New("tWK")

	// success - base flag
	flag_ := tile.Flag()
	assert.Equal(t, flag.Flag("tWK"), flag_)

	// setup
	tile = New("tWK", mockNode)

	// success - node flag
	flag_ = tile.Flag()
	assert.Equal(t, flag.Flag("mWK"), flag_)
}

func TestPop(t *testing.T) {
	// setup
	tile := New("tWK", mockNode)

	// success - node
	node := tile.Pop()
	assert.Equal(t, mockNode, node)
	assert.Empty(t, tile.Nodes)

	// success - nil
	node = tile.Pop()
	assert.Nil(t, node)
}

func TestPush(t *testing.T) {
	// setup
	tile := New("tWK")

	// success
	tile.Push(mockNode)
	assert.Equal(t, []node.Node{mockNode}, tile.Nodes)
}

func TestWall(t *testing.T) {
	// setup
	tile := New("tWK")

	// success - base wall
	okay := tile.Wall()
	assert.False(t, okay)

	// setup
	tile = New("tWK", mockNode)

	// success - node wall
	okay = tile.Wall()
	assert.True(t, okay)
}
