// Package term implements the Term type and methods.
package term

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Term is a TCell display controller.
type Term struct {
	screen tcell.Screen
}

// New returns a new Term with an initialised display.
func New() (*Term, error) {
	scrn, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	if err := scrn.Init(); err != nil {
		return nil, err
	}

	return &Term{scrn}, nil
}

// Clear erases the Term's display.
func (t *Term) Clear() {
	t.screen.Clear()
}

// Close finalises the Term's display
func (t *Term) Close() {
	t.screen.Fini()
}

// Poll blocks and returns an Event from the Term's display.
func (t *Term) Poll() tcell.Event {
	return t.screen.PollEvent()
}

// Set writes a Flag to the Term's display.
func (t *Term) Set(orig *pair.Pair, flag flag.Flag) {
	t.screen.SetContent(orig.X, orig.Y, flag.Rune(), nil, flag.Style())
}

// SetRender writes a double-spaced Flag buffer to the Term's display.
func (t *Term) SetRender(flags [][]flag.Flag) {
	wX, hY := t.screen.Size()
	mX, mY := len(flags[0])*2, len(flags)
	oX, oY := (wX-mX)/2, (hY-mY)/2

	for y, line := range flags {
		for x, flag := range line {
			t.screen.SetContent((x*2)+oX, y+oY, flag.Rune(), nil, flag.Style())
		}
	}
}

// SetPlot writes a Pair slice to the Term's display.
func (t *Term) SetPlot(pairs []*pair.Pair, flag flag.Flag) {
	for _, pair := range pairs {
		t.screen.SetContent(pair.X, pair.Y, flag.Rune(), nil, flag.Style())
	}
}

// SetString writes a string to the Term's display.
func (t *Term) SetString(orig *pair.Pair, text string, flag flag.Flag) {
	for i, char := range text {
		t.screen.SetContent(orig.X+i, orig.Y, char, nil, flag.Style())
	}
}

// Show flips the Term's display.
func (t *Term) Show() {
	t.screen.Show()
}
