// Package player implements the Player Node type and methods.
package player

import (
	"github.com/stvmln86/galah/galah/terms/flag"
)

// Player is a movable Node that represents the player in the gameworld.
type Player struct {
	flag flag.Flag
}

// New returns a new Player.
func New(flag flag.Flag) *Player {
	return &Player{flag}
}

// Flag returns the Player's drawable Flag.
func (n *Player) Flag() flag.Flag {
	return n.flag
}

// Wall returns true if the Player cannot be moved through.
func (n *Player) Wall() bool {
	return false
}
