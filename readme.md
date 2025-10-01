# Galah

**Galah** is an experimental roguelike, written in Go 1.25 by Stephen Malone.

```go

type Node interface {
	Cell() *Cell
	Name() string
}

type Cell struct {
	Rune rune
	Tone rune
}

type Tile struct {
	Base *Cell
	Slot *Node
}

type Grid struct {
	Tiles [400]*Tile
}

```
