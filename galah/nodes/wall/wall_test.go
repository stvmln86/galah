package wall

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/flag"
)

func TestNew(t *testing.T) {
	// success
	wall := New("#WK")
	assert.Equal(t, flag.Flag("#WK"), wall.flag)
}

func TestFlag(t *testing.T) {
	// setup
	wall := New("#WK")

	// success
	flag_ := wall.Flag()
	assert.Equal(t, flag.Flag("#WK"), flag_)
}

func TestWall(t *testing.T) {
	// setup
	wall := New("#WK")

	// success
	okay := wall.Wall()
	assert.True(t, okay)
}
