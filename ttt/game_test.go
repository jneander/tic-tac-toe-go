package ttt

import "github.com/sdegutis/go.assert"
import "testing"

func TestGameBoard( t *testing.T ) {
  var game Game = NewGame()

  t.Log( "Board() returns the game's board instance" )
  var board *Board = game.Board()
  assert.Equals( t, game.Board(), board )
}

func TestGameIsOver( t *testing.T ) {
  var game Game = NewGame()
  var board = game.Board()

  t.Log( "New Game is not over" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with two-in-a-row is not over" )
  AddMarks( board, []int{ 4, 8 }, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"X\" is over" )
  board.Mark( 0, "X" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"O\" is over" )
  board.Reset()
  AddMarks( board, []int{ 2, 4, 6 }, "O" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row mismatched is not over" )
  board.Mark( 2, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with nearly-full, non-winning board is not over" )
  board.Reset()
  AddMarks( board, []int{ 0, 1, 4, 5 }, "X" )
  AddMarks( board, []int{ 2, 3, 7, 8 }, "O" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with full, non-winning board is over" )
  board.Mark( 6, "X" )
  assert.True( t, game.IsOver() )
}

func TestGameIsValidMove( t *testing.T ) {
  var game  = NewGame()
  var board = game.Board()

  t.Log( "IsValidMove() returns true if the selected space is blank" )
  assert.True( t, game.IsValidMove( 1 ) )
  assert.True( t, game.IsValidMove( 2 ) )

  t.Log( "IsValidMove() returns true if the selected space is blank" )
  board.Mark( 1, "X" )
  board.Mark( 2, "O" )
  assert.False( t, game.IsValidMove( 1 ) )
  assert.False( t, game.IsValidMove( 2 ) )
}

func TestGameApplyMove( t *testing.T ) {
  var game  = NewGame()
  var board = game.Board()

  t.Log( "ApplyMove() applies the selected mark to the board" )
  game.ApplyMove( 0, "X" )
  game.ApplyMove( 1, "O" )
  assert.Equals( t, board.Spaces()[0], "X" )
  assert.Equals( t, board.Spaces()[1], "O" )

  t.Log( "ApplyMove() rejects moves for non-blank spaces" )
  game.ApplyMove( 1, "X" )
  game.ApplyMove( 0, "O" )
  assert.Equals( t, board.Spaces()[0], "X" )
  assert.Equals( t, board.Spaces()[1], "O" )
}

func TestGameReset( t *testing.T ) {
  var game  = NewGame()
  var board = game.Board()

  t.Log( "Reset() clears the board" )
  game.ApplyMove( 0, "X" )
  game.ApplyMove( 1, "O" )
  game.Reset()
  assert.Equals( t, board.Spaces()[0], board.Blank() )
  assert.Equals( t, board.Spaces()[1], board.Blank() )
}

func AddMarks( b *Board, set []int, mark string ) {
  for _,p := range set {
    b.Mark( p, mark )
  }
}
