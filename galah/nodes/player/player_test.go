package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/terms/flag"
)

func TestNew(t *testing.T) {
	// success
	plyr := New("@WK")
	assert.Equal(t, flag.Flag("@WK"), plyr.flag)
}

func TestFlag(t *testing.T) {
	// setup
	plyr := New("@WK")

	// success
	flag_ := plyr.Flag()
	assert.Equal(t, flag.Flag("@WK"), flag_)
}

func TestWall(t *testing.T) {
	// setup
	plyr := New("mWK")

	// success
	okay := plyr.Wall()
	assert.True(t, okay)
}
