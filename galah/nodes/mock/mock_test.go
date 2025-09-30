package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/flag"
)

func TestNew(t *testing.T) {
	// success
	mock := New("mWK", true)
	assert.Equal(t, flag.Flag("mWK"), mock.flag)
	assert.Equal(t, true, mock.wall)
}

func TestFlag(t *testing.T) {
	// setup
	mock := New("mWK", true)

	// success
	flag_ := mock.Flag()
	assert.Equal(t, flag.Flag("mWK"), flag_)
}

func TestWall(t *testing.T) {
	// setup
	mock := New("mWK", true)

	// success
	okay := mock.Wall()
	assert.True(t, okay)
}
