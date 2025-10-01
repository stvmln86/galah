package cell

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

var (
	mockCell = New('A', 'B')
)

func TestNew(t *testing.T) {
	// success
	cell := New('A', 'B')
	assert.Equal(t, 'A', cell.rune)
	assert.Equal(t, 'B', cell.tone)
}

func TestRune(t *testing.T) {
	// success
	rune := mockCell.Rune()
	assert.Equal(t, 'A', rune)
}

func TestStyle(t *testing.T) {
	// success
	styl := mockCell.Style()
	fore, back, attr := styl.Decompose()
	assert.Equal(t, tcell.ColorBlue, fore)
	assert.Equal(t, tcell.ColorDefault, back)
	assert.Equal(t, tcell.AttrNone, attr)
}
