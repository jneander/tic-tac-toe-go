package ttt

import "github.com/stretchrcom/testify/assert"
import "testing"

func TestGame_Board( t *testing.T ) {
  var game Game = NewGame()

  t.Log( "#Board returns the game's board instance" )
  var board *Board = game.Board()
  assert.Equal( t, game.Board(), board )
}

func TestGame_IsOver( t *testing.T ) {
  var game Game = NewGame()
  var board = game.Board()

  t.Log( "New Game is not over" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with two-in-a-row is not over" )
  AddMarks( board, "X", 4, 8 )
  assert.False( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"X\" is over" )
  board.Mark( 0, "X" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"O\" is over" )
  board.Reset()
  AddMarks( board, "O", 2, 4, 6 )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row mismatched is not over" )
  board.Mark( 2, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with nearly-full, non-winning board is not over" )
  board.Reset()
  AddMarks( board, "X", 0, 1, 4, 5 )
  AddMarks( board, "O", 2, 3, 7, 8 )
  assert.False( t, game.IsOver() )

  t.Log( "Game with full, non-winning board is over" )
  board.Mark( 6, "X" )
  assert.True( t, game.IsOver() )
}

func TestGame_IsValidMove( t *testing.T ) {
  var game  = NewGame()
  var board = game.Board()

  t.Log( "#IsValidMove returns true if the selected space is blank" )
  assert.True( t, game.IsValidMove( 1 ) )
  assert.True( t, game.IsValidMove( 2 ) )

  t.Log( "#IsValidMove returns true if the selected space is blank" )
  board.Mark( 1, "X" )
  board.Mark( 2, "O" )
  assert.False( t, game.IsValidMove( 1 ) )
  assert.False( t, game.IsValidMove( 2 ) )

  t.Log( "#IsValidMove returns false if the provided index is out of range" )
  assert.False( t, game.IsValidMove( -1 ) )
  assert.False( t, game.IsValidMove( 9 ) )
}

func TestGame_ApplyMove( t *testing.T ) {
  var game  = NewGame()
  var board = game.Board()

  t.Log( "#ApplyMove applies the selected mark to the board" )
  game.ApplyMove( 0, "X" )
  game.ApplyMove( 1, "O" )
  assert.Equal( t, board.Spaces()[0], "X" )
  assert.Equal( t, board.Spaces()[1], "O" )

  t.Log( "#ApplyMove rejects moves for non-blank spaces" )
  game.ApplyMove( 1, "X" )
  game.ApplyMove( 0, "O" )
  assert.Equal( t, board.Spaces()[0], "X" )
  assert.Equal( t, board.Spaces()[1], "O" )
}

func TestGame_Reset( t *testing.T ) {
  var game  = NewGame()
  var board = game.Board()

  t.Log( "#Reset clears the board" )
  game.ApplyMove( 0, "X" )
  game.ApplyMove( 1, "O" )
  game.Reset()
  assert.Equal( t, board.Spaces()[0], board.Blank() )
  assert.Equal( t, board.Spaces()[1], board.Blank() )
}
