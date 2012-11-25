package ttt

import "github.com/sdegutis/go.assert"
import "testing"
import "bytes"

var in   bytes.Buffer
var out  bytes.Buffer
var ui UI = ConsoleUI{ &in, &out }

func TestNewConsoleUI( t *testing.T ) {
  var ui *ConsoleUI = NewConsoleUI( &in, &out )
  assert.Equals( t, ui.in, &in )
  assert.Equals( t, ui.out, &out )
}

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
  expected := "\n     _|_|_\n     _|_|_\n     _|_|_\n\n"
  assert.Equals( t, ReadInput( &out ), expected )

  t.Log( "DisplayBoard() prints a board with marks" )
  ints, marks := []int{ 4, 5, 6, 8 }, []string{ "X", "O", "O", "X" }
  for i := range ints { board.Mark( ints[i], marks[i] ) }
  ui.DisplayBoard( board )
  expected = "\n     _|_|_\n     _|X|O\n     O|_|X\n\n"
  assert.Equals( t, ReadInput( &out ), expected )
}

func TestConsoleUiPromptPlayerMove( t *testing.T ) {
  t.Log( "PromptPlayerMove() prints a prompt" )
  SetInputString( &in, "4\n" )
  ui.PromptPlayerMove()
  expected := "Please enter the space for your mark: "
  assert.Equals( t, ReadInput( &out ), expected )

  t.Log( "PromptPlayerMove() reprints the prompt after invalid input" )
  SetInputString( &in, "\n4" )
  ui.PromptPlayerMove()
  expected = "Please enter the space for your mark: "
  expected += expected
  assert.Equals( t, ReadInput( &out ), expected )

  t.Log( "PromptPlayerMove() returns the user's input" )
  SetInputString( &in, "4\n5\n" )
  assert.Equals( t, ui.PromptPlayerMove(), 4 )
  assert.Equals( t, ui.PromptPlayerMove(), 5 )

  t.Log( "PromptPlayerMove() rejects input not found in optional filter list" )
  SetInputString( &in, "3\n5\n7" )
  assert.Equals( t, ui.PromptPlayerMove( 4, 5, 6 ), 5 )

  t.Log( "PromptPlayerMove() rejects invalid input" )
  SetInputString( &in, "\ninvalid\n5" )
  assert.Equals( t, ui.PromptPlayerMove(), 5 )
}

func TestReadLine( t *testing.T ) {
  t.Log( "ReadLine() reads input up until newline" )
  buffer := bytes.NewBuffer( []byte( "test\nvalue" ) )
  assert.Equals( t, ReadLine( buffer ), "test" )

  t.Log( "ReadLine() reads input up until end of reader buffer" )
  assert.Equals( t, ReadLine( buffer ), "value" )
  assert.Equals( t, ReadLine( buffer ), "" )
}

func SetInputString( input *bytes.Buffer, data string ) {
  input.Reset();
  input.WriteString( data )
}
