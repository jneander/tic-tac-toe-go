package ttt

import (
	"bytes"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

func TestConsolePlayer_Move(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer
	var console = NewConsole(&in, &out)
	var player Player = NewConsolePlayer(console)
	var board = NewBoard()

	t.Log("returns the input from the Console")
	SetInputs(&in, MovesAsInput(2, 4, 6)...)
	assert.Equal(t, player.Move(*board), 2)
	assert.Equal(t, player.Move(*board), 4)
	assert.Equal(t, player.Move(*board), 6)

	t.Log("restricts responses to unmarked spaces")
	SetInputs(&in, MovesAsInput(3, 4, 5)...)
	AddMarks(board, "X", 2, 3, 6)
	assert.Equal(t, player.Move(*board), 4)
	assert.Equal(t, player.Move(*board), 5)
}

func TestConsolePlayer_Mark(t *testing.T) {
	var player = NewConsolePlayer(nil)

	t.Log("implements GetMark and SetMark")
	player.SetMark("X")
	assert.Equal(t, player.GetMark(), "X")
	player.SetMark("O")
	assert.Equal(t, player.GetMark(), "O")
}
