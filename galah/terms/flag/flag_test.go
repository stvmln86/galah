package flag

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestBack(t *testing.T) {
	// setup
	flag := Flag("#WK")

	// success
	back := flag.Back()
	assert.Equal(t, tcell.ColorBlack, back)
}

func TestFore(t *testing.T) {
	// setup
	flag := Flag("#WK")

	// success
	back := flag.Fore()
	assert.Equal(t, tcell.ColorWhite, back)
}

func TestRune(t *testing.T) {
	// setup
	flag := Flag("#WK")

	// success
	rune := flag.Rune()
	assert.Equal(t, '#', rune)
}

func TestStyle(t *testing.T) {
	// setup
	flag := Flag("#WK")
	want := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// success
	styl := flag.Style()
	assert.Equal(t, want, styl)
}
