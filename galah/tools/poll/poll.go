// Package poll implements terminal event polling functions and constants.
package poll

import (
	"github.com/gdamore/tcell/v2"
)

// binds is a map of all defined special key presses.
var binds = map[tcell.Key]string{
	// general keys
	tcell.KeyBackspace: "backspace",
	tcell.KeyEnter:     "enter",
	tcell.KeyEscape:    "escape",
	tcell.KeyTab:       "tab",

	// arrow keys
	tcell.KeyUp:    "up",
	tcell.KeyDown:  "down",
	tcell.KeyLeft:  "left",
	tcell.KeyRight: "right",
}

// Poll returns an event string from a Screen event.
func Poll(scrn tcell.Screen) string {
	switch evnt := scrn.PollEvent().(type) {
	case *tcell.EventKey:
		if name, ok := binds[evnt.Key()]; ok {
			return name
		} else {
			return string(evnt.Rune())
		}

	case *tcell.EventResize:
		return "resize"

	default:
		return "unknown"
	}
}
