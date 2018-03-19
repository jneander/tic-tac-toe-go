package ttt

import (
	"bytes"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

var in bytes.Buffer
var out bytes.Buffer
var ui UI = NewConsole(&in, &out)

func TestConsole_PromptMainMenu(t *testing.T) {
	t.Log("#PromptMainMenu displays a prompt message")
	SetInputs(&in, "3")
	expected := "\nWelcome to Tic Tac Toe in Go!\n" +
		"1) Player Goes First\n" +
		"2) Computer Goes First\n" +
		"3) Exit\n\n" +
		"Please enter your choice: "
	ui.PromptMainMenu()
	actual := ReadInput(&out)
	assert.Equal(t, actual, expected)

	t.Log("#PromptMainMenu accepts only options listed")
	SetInputs(&in, "9", "25", "3")
	assert.Equal(t, ui.PromptMainMenu(), EXIT_GAME)

	t.Log("#PromptMainMenu returns constants based on input")
	SetInputs(&in, "1", "2", "3")
	assert.Equal(t, ui.PromptMainMenu(), PLAYER_FIRST)
	assert.Equal(t, ui.PromptMainMenu(), COMPUTER_FIRST)
	assert.Equal(t, ui.PromptMainMenu(), EXIT_GAME)
}

func TestConsole_DisplayBoard(t *testing.T) {
	var board = NewBoard()

	t.Log("#DisplayBoard prints an empty board")
	out.Reset()
	ui.DisplayBoard(board)
	expected := "\n     _|_|_\n     _|_|_\n     _|_|_\n\n"
	assert.Equal(t, ReadInput(&out), expected)

	t.Log("#DisplayBoard prints a board with marks")
	out.Reset()
	AddMarks(board, "X", 4, 8)
	AddMarks(board, "O", 5, 6)
	ui.DisplayBoard(board)
	expected = "\n     _|_|_\n     _|X|O\n     O|_|X\n\n"
	assert.Equal(t, ReadInput(&out), expected)
}

func TestConsole_DisplayAvailableSpaces(t *testing.T) {
	var board = NewBoard()

	t.Log("#DisplayAvailableSpaces prints board and all spaces")
	ui.DisplayAvailableSpaces(board)
	expected := "\n     _|_|_     1 2 3\n" +
		"     _|_|_     4 5 6\n" +
		"     _|_|_     7 8 9\n\n"
	assert.Equal(t, ReadInput(&out), expected)

	t.Log("#DisplayAvailableSpaces prints board and all spaces")
	AddMarks(board, "X", 4, 8)
	AddMarks(board, "O", 5, 6)
	ui.DisplayAvailableSpaces(board)
	expected = "\n     _|_|_     1 2 3\n" +
		"     _|X|O     4    \n" +
		"     O|_|X       8  \n\n"
	assert.Equal(t, ReadInput(&out), expected)
}

func TestConsole_DisplayGameOver(t *testing.T) {
	var game = NewGame()
	var board = game.Board()

	t.Log("#DisplayGameOver reports 'Player X' win")
	AddMarks(board, "X", 0, 1, 2)
	ui.DisplayGameOver(game)
	expected := "Player X is the winner!\n\n"
	assert.Equal(t, ReadInput(&out), expected)

	t.Log("#DisplayGameOver reports 'Player X' win")
	AddMarks(board, "O", 1, 4, 7)
	ui.DisplayGameOver(game)
	expected = "Player O is the winner!\n\n"
	assert.Equal(t, ReadInput(&out), expected)

	t.Log("#DisplayGameOver informs of draw game")
	board.Reset()
	ui.DisplayGameOver(game)
	expected = "The game has ended in a draw!\n\n"
	assert.Equal(t, ReadInput(&out), expected)
}

func TestConsole_PromptPlayerMove(t *testing.T) {
	t.Log("#PromptPlayerMove prints a prompt")
	SetInputs(&in, "4")
	ui.PromptPlayerMove()
	expected := "Please enter the space for your mark: "
	assert.Equal(t, ReadInput(&out), expected)

	t.Log("#PromptPlayerMove reprints the prompt after invalid input")
	SetInputs(&in, "", "5")
	ui.PromptPlayerMove()
	expected = "Please enter the space for your mark: "
	expected += expected
	assert.Equal(t, ReadInput(&out), expected)

	t.Log("#PromptPlayerMove returns the user's input")
	SetInputs(&in, MovesAsInput(5, 6)...)
	assert.Equal(t, ui.PromptPlayerMove(), 5)
	assert.Equal(t, ui.PromptPlayerMove(), 6)

	t.Log("#PromptPlayerMove rejects input not found in optional filter list")
	SetInputs(&in, MovesAsInput(3, 5, 7)...)
	assert.Equal(t, ui.PromptPlayerMove(4, 5, 6), 5)

	t.Log("#PromptPlayerMove rejects invalid input")
	SetInputs(&in, "", "invalid", "6")
	assert.Equal(t, ui.PromptPlayerMove(), 6-1)
}

func Test_ReadLine(t *testing.T) {
	t.Log("ReadLine() reads input up until newline")
	buffer := bytes.NewBuffer([]byte("test\nvalue"))
	assert.Equal(t, ReadLine(buffer), "test")

	t.Log("ReadLine() reads input up until end of reader buffer")
	assert.Equal(t, ReadLine(buffer), "value")
	assert.Equal(t, ReadLine(buffer), "")
}
