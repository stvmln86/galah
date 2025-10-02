// Package icon implements the Icon interface and methods.
package icon

import "github.com/gdamore/tcell/v2"

// Icon is an interface for terminal display characters.
type Icon interface {
	// Char returns the Icon's display character.
	Char() rune

	// Tone returns the Icon's foreground colour.
	Tone() tcell.Color
}

// Style returns a terminal style from an Icon.
func Style(icon Icon) tcell.Style {
	return tcell.StyleDefault.Foreground(icon.Tone())
}
