package tictactoe

import "github.com/sdegutis/go.assert"
import "testing"
import "reflect"

var BLANK = NewBoard().Blank()

func TestBoardInitialization( t *testing.T ) {
  board := NewBoard()
  spaces := board.Spaces()

  t.Log( "Board has nine spaces" )
  assert.Equals( t, len( spaces ), 9 )

  t.Logf( "Board is initialized with \"%s\" characters", BLANK )
  for _, v := range spaces {
    assert.Equals( t, v, BLANK )
  }
}

func TestBoardConstants( t *testing.T ) {
  board := NewBoard()

  t.Log( "Board.Blank() returns the mark representing blanks" )
  assert.Equals( t, board.Blank(), BLANK )
}

func TestBoardProtection( t *testing.T ) {
  board := NewBoard()

  t.Log( "Board.Spaces() returns a copy of the spaces" )
  p1 := reflect.ValueOf( board.Spaces() )
  p2 := reflect.ValueOf( board.Spaces() )
  assert.NotEquals( t, p1.Pointer(), p2.Pointer() )

  t.Log( "Board spaces array cannot be directly modified" )
  spaces := board.Spaces()
  spaces[0] = "O"
  assert.NotEquals( t, board.Spaces()[0], "O" )
}

func TestBoardMarking( t *testing.T ) {
  board := NewBoard()

  t.Log( "Board.Mark() assigns a mark to the board at a given index" )
  board.Mark( 4, "X" )
  assert.Equals( t, board.Spaces()[4], "X" )

  t.Log( "Board.Mark() reassigns a mark to the board at a given index" )
  board.Mark( 4, "O" )
  assert.Equals( t, board.Spaces()[4], "O" )

  t.Log( "Board.Mark() ignores indices out of range" )
  var spaces = board.Spaces()
  board.Mark( -1, "X" )
  assert.DeepEquals( t, board.Spaces(), spaces )
  board.Mark( 9, "X" )
  assert.DeepEquals( t, board.Spaces(), spaces )
}

func TestBoardReset( t *testing.T ) {
  board := NewBoard()
  spaces := board.Spaces()
  board.Mark( 0, "X" )
  board.Mark( 8, "O" )

  t.Log( "Board.Reset() resets spaces to initialized state" )
  board.Reset()
  assert.DeepEquals( t, board.Spaces(), spaces )
}
