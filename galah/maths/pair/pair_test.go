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
