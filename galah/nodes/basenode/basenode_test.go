package basenode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/icons/baseicon"
	"github.com/stvmln86/galah/galah/tools/tone"
)

var mockIcon = baseicon.New('i', tone.Blue)

func TestNew(t *testing.T) {
	// success
	node := New(mockIcon, "test", true)
	assert.Equal(t, mockIcon, node.icon)
	assert.Equal(t, "test", node.name)
	assert.Equal(t, true, node.open)
}

func TestNewFunc(t *testing.T) {
	// success
	nfun := NewFunc(mockIcon, "test", true)
	node := nfun().(*BaseNode)
	assert.Equal(t, mockIcon, node.icon)
	assert.Equal(t, "test", node.name)
	assert.Equal(t, true, node.open)
}

func TestIcon(t *testing.T) {
	// setup
	node := New(mockIcon, "test", true)

	// success
	icon := node.Icon()
	assert.Equal(t, mockIcon, icon)
}

func TestName(t *testing.T) {
	// setup
	node := New(mockIcon, "test", true)

	// success
	name := node.Name()
	assert.Equal(t, "test", name)
}

func TestOpen(t *testing.T) {
	// setup
	node := New(mockIcon, "test", true)

	// success
	open := node.Open()
	assert.Equal(t, true, open)
}
