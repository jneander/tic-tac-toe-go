package tictactoe

import "github.com/sdegutis/go.assert"
import "testing"
import "bytes"

var in   bytes.Buffer
var out  bytes.Buffer
var ui UI = ConsoleUI{ &in, &out }

func TestConsoleUiPromptMainMenu( t *testing.T ) {
  in.WriteString( "any key...\n" )
  expected := "Welcome to Tic Tac Toe in Go!\n" +
              "Press any key to exit... "

  ui.PromptMainMenu()
  actual := ReadInput( &out )
  assert.Equals( t, actual, expected )
}

func TestConsoleUiDisplayBoard( t *testing.T ) {
  board := NewBoard()

  t.Log( "DisplayBoard() prints an empty board" )
  ui.DisplayBoard( board )
  expected := "_|_|_\n_|_|_\n_|_|_\n"
  assert.Equals( t, ReadInput( &out ), expected )

  t.Log( "DisplayBoard() prints a board with marks" )
  ints, marks := []int{ 4, 5, 6, 8 }, []string{ "X", "O", "O", "X" }
  for i := range ints { board.Mark( ints[i], marks[i] ) }
  ui.DisplayBoard( board )
  expected = "_|_|_\n_|X|O\nO|_|X\n"
  assert.Equals( t, ReadInput( &out ), expected )
}
