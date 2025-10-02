// Package baseicon implements the BaseIcon type and methods.
package baseicon

import (
	"github.com/stvmln86/galah/galah/tools/tone"
)

// BaseIcon is a basic implementation of an Icon.
type BaseIcon struct {
	char rune
	tone tone.Colour
}

// New returns a new BaseIcon.
func New(char rune, tone tone.Colour) *BaseIcon {
	return &BaseIcon{char, tone}
}

// Char returns the Icon's display character.
func (i *BaseIcon) Char() rune {
	return i.char
}

// Tone returns the Icon's foreground colour.
func (i *BaseIcon) Tone() tone.Colour {
	return i.tone
}
