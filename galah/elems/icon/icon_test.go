package icon

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/galah/galah/tools/tone"
)

type mockIcon struct{}

func (i *mockIcon) Char() rune        { return 'm' }
func (i *mockIcon) Tone() tcell.Color { return tone.Blue }

func TestStyle(t *testing.T) {
	// setup
	icon := new(mockIcon)

	// success
	styl := Style(icon)
	fore, back, attr := styl.Decompose()
	assert.Equal(t, tcell.ColorBlue, fore)
	assert.Equal(t, tcell.ColorDefault, back)
	assert.Equal(t, tcell.AttrNone, attr)
}
