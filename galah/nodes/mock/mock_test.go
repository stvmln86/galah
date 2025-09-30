package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/flag"
)

func TestNew(t *testing.T) {
	// success
	mock := New("#WK", true)
	assert.Equal(t, flag.Flag("#WK"), mock.flag)
	assert.Equal(t, true, mock.wall)
}

func TestFlag(t *testing.T) {
	// setup
	mock := New("#WK", true)

	// success
	flag_ := mock.Flag()
	assert.Equal(t, flag.Flag("#WK"), flag_)
}

func TestWall(t *testing.T) {
	// setup
	mock := New("#WK", true)

	// success
	okay := mock.Wall()
	assert.True(t, okay)
}
