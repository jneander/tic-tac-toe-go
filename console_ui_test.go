package tictactoe

import "github.com/sdegutis/go.assert"
import "testing"
import "bytes"

var in   bytes.Buffer
var out  bytes.Buffer
var ui = Console{ &in, &out }

func TestConsoleUiPromptMainMenu( t *testing.T ) {
  in.WriteString( "any key...\n" )
  expected := "Welcome to Tic Tac Toe in Go!\n" +
              "Press any key to exit... "

  ui.PromptMainMenu()
  actual := ReadInput( &out )
  assert.Equals( t, actual, expected )
}
