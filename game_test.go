package tictactoe

import "github.com/sdegutis/go.assert"
import "testing"

func TestGameOver( t *testing.T ) {
  var game Game = NewGame()

  t.Log( "New Game is not over" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with two-in-a-row is not over" )
  addMarks( game.Board, []int{ 4, 8 }, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"X\" is over" )
  game.Board.Mark( 0, "X" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"O\" is over" )
  game.Board.Reset()
  addMarks( game.Board, []int{ 2, 4, 6 }, "O" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row mismatched is not over" )
  game.Board.Mark( 2, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with nearly-full, non-winning board is not over" )
  game.Board.Reset()
  addMarks( game.Board, []int{ 0, 1, 4, 5 }, "X" )
  addMarks( game.Board, []int{ 2, 3, 7, 8 }, "O" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with full, non-winning board is over" )
  game.Board.Mark( 6, "X" )
  assert.True( t, game.IsOver() )
}

func TestGameIsValidMove( t *testing.T ) {
  var game  = NewGame()
  var board = &game.Board

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
  var board = &game.Board

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

func addMarks( b *Board, set []int, mark string ) {
  for _,p := range set {
    b.Mark( p, mark )
  }
}
