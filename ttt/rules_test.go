package ttt

import "github.com/stretchrcom/testify/assert"
import "testing"

func TestRules_MarkHasWinningSolution( t *testing.T ) {
  var rules *Rules = NewRules()
  var board *Board = NewBoard()

  t.Log( "returns false when board is empty" )
  assert.False( t, rules.MarkHasWinningSolution( board, "X" ) )
  assert.False( t, rules.MarkHasWinningSolution( board, "O" ) )

  t.Log( "returns true for three in a row of the given mark" )
  AddMarks( board, "X", 3, 4, 5 )
  assert.True( t, rules.MarkHasWinningSolution( board, "X" ) )
  assert.False( t, rules.MarkHasWinningSolution( board, "O" ) )
  AddMarks( board, "O", 1, 4, 7 )
  assert.True( t, rules.MarkHasWinningSolution( board, "O" ) )
  assert.False( t, rules.MarkHasWinningSolution( board, "X" ) )

  t.Log( "returns false if no win and board is full" )
  AddMarks( board, "X", 0, 1, 5, 6, 7 )
  AddMarks( board, "O", 2, 3, 4, 8 )
  assert.False( t, rules.MarkHasWinningSolution( board, "X" ) )
  assert.False( t, rules.MarkHasWinningSolution( board, "O" ) )
}
