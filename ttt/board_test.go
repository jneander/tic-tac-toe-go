package ttt

import "github.com/stretchrcom/testify/assert"
import sassert "github.com/sdegutis/go.assert"
import "testing"
import "reflect"

var BLANK = NewBoard().Blank()

func TestBoardInitialization( t *testing.T ) {
  board := NewBoard()
  spaces := board.Spaces()

  t.Log( "Board has nine spaces" )
  assert.Equal( t, len( spaces ), 9 )

  t.Logf( "Board is initialized with \"%s\" characters", BLANK )
  for _, v := range spaces {
    assert.Equal( t, v, BLANK )
  }
}

func TestBoardConstants( t *testing.T ) {
  board := NewBoard()

  t.Log( "Board.Blank() returns the mark representing blanks" )
  assert.Equal( t, board.Blank(), BLANK )
}

func TestBoardProtection( t *testing.T ) {
  board := NewBoard()

  t.Log( "Board.Spaces() returns a copy of the spaces" )
  p1 := reflect.ValueOf( board.Spaces() )
  p2 := reflect.ValueOf( board.Spaces() )
  assert.NotEqual( t, p1.Pointer(), p2.Pointer() )

  t.Log( "Board spaces array cannot be directly modified" )
  spaces := board.Spaces()
  spaces[0] = "O"
  assert.NotEqual( t, board.Spaces()[0], "O" )
}

func TestBoardMarking( t *testing.T ) {
  board := NewBoard()

  t.Log( "Board.Mark() assigns a mark to the board at a given index" )
  board.Mark( 4, "X" )
  assert.Equal( t, board.Spaces()[4], "X" )

  t.Log( "Board.Mark() reassigns a mark to the board at a given index" )
  board.Mark( 4, "O" )
  assert.Equal( t, board.Spaces()[4], "O" )

  t.Log( "Board.Mark() ignores indices out of range" )
  var spaces = board.Spaces()
  board.Mark( -1, "X" )
  sassert.DeepEquals( t, board.Spaces(), spaces )
  board.Mark( 9, "X" )
  sassert.DeepEquals( t, board.Spaces(), spaces )
}

func TestBoardReset( t *testing.T ) {
  board := NewBoard()
  spaces := board.Spaces()
  board.Mark( 0, "X" )
  board.Mark( 8, "O" )

  t.Log( "Board.Reset() resets spaces to initialized state" )
  board.Reset()
  sassert.DeepEquals( t, board.Spaces(), spaces )
}

func TestBoardSpacesWithMark( t *testing.T ) {
  board := NewBoard()
  sassert.DeepEquals( t, board.SpacesWithMark( "X" ), []int{} )
  AddMarks( board, "X", 1, 3, 5 )
  AddMarks( board, "O", 2, 4, 8 )
  sassert.DeepEquals( t, board.SpacesWithMark( "X" ), []int{ 1, 3, 5 } )
  sassert.DeepEquals( t, board.SpacesWithMark( "O" ), []int{ 2, 4, 8 } )
  sassert.DeepEquals( t, board.SpacesWithMark( board.Blank() ), []int{ 0, 6, 7 } )
}
