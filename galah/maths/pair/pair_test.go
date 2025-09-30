package pair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	pair := New(1, 2)
	assert.Equal(t, 1, pair.X)
	assert.Equal(t, 2, pair.Y)
}

func TestAdd(t *testing.T) {
	// setup
	pair := New(1, 2)

	// success
	pair2 := pair.Add(pair)
	assert.Equal(t, 2, pair2.X)
	assert.Equal(t, 4, pair2.Y)
}

func TestAddXY(t *testing.T) {
	// setup
	pair := New(1, 2)

	// success
	pair2 := pair.AddXY(1, 2)
	assert.Equal(t, 2, pair2.X)
	assert.Equal(t, 4, pair2.Y)
}
