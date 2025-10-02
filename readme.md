# Galah

**Galah** is an experimental roguelike, written in Go 1.25 by Stephen Malone.

```

galah/
  - games/
    - grid/grid.go · game grid
    - icon/icon.go · Icon { Char() rune, Tone() rune }
    - node/node.go · Node { Icon() Icon, Name() string, Open() bool }

  - icons/
    - baseicon/baseicon.go

  - nodes/
    - basenode/basenode.go

  - tools/
    - draw/draw.go · tcell drawing functions
    - plot/plot.go · math functions
    - tone/tone.go · colour constants
```
