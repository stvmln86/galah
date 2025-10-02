package baseicon

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/tools/tone"
)

func TestNew(t *testing.T) {
	// success
	icon := New('i', tone.Blue)
	assert.Equal(t, 'i', icon.char)
	assert.Equal(t, tone.Blue, icon.tone)
}

func TestChar(t *testing.T) {
	// setup
	icon := New('i', tone.Blue)

	// success
	char := icon.Char()
	assert.Equal(t, 'i', char)
}

func TestTone(t *testing.T) {
	// setup
	icon := New('i', tone.Blue)

	// success
	tone_ := icon.Tone()
	assert.Equal(t, tone.Blue, tone_)
}
