package poll

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	// setup
	scrn := tcell.NewSimulationScreen("utf-8")
	scrn.Init()
	scrn.InjectKey(tcell.KeyEnter, ' ', 0)
	defer scrn.Fini()

	// success
	evnt := Poll(scrn)
	assert.Equal(t, "enter", evnt)
}
