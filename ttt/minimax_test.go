package ttt

import sassert "github.com/sdegutis/go.assert"
import "github.com/stretchrcom/testify/assert"
import "testing"

var testDeepRecursion = false

func TestMinimax_ScoreAvailableMoves( t *testing.T ) {
  var minimax *Minimax = NewMinimax()
  var board = NewBoard()
  var min, max = "O", "X"
  minimax.SetMinMaxMarks( min, max )

  t.Log( "DepthLimit is initialized to 5" )
  assert.Equal( t, minimax.DepthLimit, 5 )

  // set marks with possible win, loss, or block
  AddMarks( board, min, 0, 3 )
  AddMarks( board, max, 2, 5 )

  // without recursion
  minimax.DepthLimit = 0

  t.Log( "returns the board scores for the current depth" )
  expected := map[int]int{ 1:0, 4:0, 6:0, 7:0, 8:1 }
  actual,_ := minimax.ScoreAvailableMoves( board, max )
  sassert.DeepEquals( t, actual, expected )

  // with one level of recursion
  minimax.DepthLimit = 1

  t.Log( "evaluates the board through to the next depth" )
  expected = map[int]int{ 1:-1, 4:-1, 6:0, 7:-1, 8:1 }
  actual,_ = minimax.ScoreAvailableMoves( board, max )
  sassert.DeepEquals( t, actual, expected )

  // set marks with guaranteed win in two turns
  board.Reset()
  AddMarks( board, min, 1, 5 )
  AddMarks( board, max, 0, 7 )

  // without recursion
  minimax.DepthLimit = 0

  t.Log( "returns the scores for each available space" )
  expected = map[int]int{ 2:0, 3:0, 4:0, 6:0, 8:0 }
  actual,_ = minimax.ScoreAvailableMoves( board, max )
  sassert.DeepEquals( t, actual, expected )

  // with one level of recursion
  minimax.DepthLimit = 1

  t.Log( "returns the scores after one level of recursion" )
  expected = map[int]int{ 2:0, 3:0, 4:0, 6:0, 8:0 }
  actual,_ = minimax.ScoreAvailableMoves( board, max )
  sassert.DeepEquals( t, actual, expected )

  // with two levels of recursion
  minimax.DepthLimit = 2

  t.Log( "returns the scores after two levels of recursion" )
  expected = map[int]int{ 2:0, 3:0, 4:0, 6:1, 8:1 }
  actual,final := minimax.ScoreAvailableMoves( board, max )
  sassert.DeepEquals( t, actual, expected )

  t.Log( "returns false if any scores are not final" )
  assert.False( t, final )

  t.Log( "returns true if all scores are final" )
  board.Reset()
  AddMarks( board, max, 1, 7 )
  AddMarks( board, min, 3, 5 )
  actual,final = minimax.ScoreAvailableMoves( board, max )
  assert.True( t, final )

  // set marks with guaranteed win in two turns
  board.Reset()
  AddMarks( board, min, 1, 5 )
  AddMarks( board, max, 0, 6, 7 )

  // without recursion
  minimax.DepthLimit = 1

  t.Log( "returns results of calls to #Score" )
  expected = map[int]int{ 2:1, 3:1, 4:1, 8:1 }
  actual,_ = minimax.ScoreAvailableMoves( board, min )
  sassert.DeepEquals( t, actual, expected )

  // DEEP RECURSION

  if testDeepRecursion {
    // minimum starting recursion (the default)
    minimax.DepthLimit = 5

    // opponent takes the center
    board.Reset()
    AddMarks( board, min, 4 )

    t.Log( "returns -1 for all edge spaces" )
    expected = map[int]int{ 0:0, 1:-1, 2:0, 3:-1, 5:-1, 6:0, 7:-1, 8:0 }
    actual,_ = minimax.ScoreAvailableMoves( board, max )
    sassert.DeepEquals( t, actual, expected )
  }
}

func TestMinimax_Score( t *testing.T ) {
  var minimax = NewMinimax()
  var board = NewBoard()
  var min, max = "X", "O"
  minimax.SetMinMaxMarks( min, max )

  // without recursion
  minimax.DepthLimit = 0

  t.Log( "#Score returns 1, true if 'max mark' won" )
  AddMarks( board, max, 3, 4, 5 )
  score,final := minimax.Score( board, max )
  assert.Equal( t, score, 1 )
  assert.True( t, final )

  t.Log( "#Score returns -1, true if 'min mark' won" )
  AddMarks( board, min, 1, 4, 7 )
  score,final = minimax.Score( board, max )
  assert.Equal( t, score, -1 )
  assert.True( t, final )

  t.Log( "#Score returns 0, true for no-win, full board" )
  AddMarks( board, max, 0, 1, 5, 6, 7 )
  AddMarks( board, min, 2, 3, 4, 8 )
  score,final = minimax.Score( board, max )
  assert.Equal( t, score, 0 )
  assert.True( t, final )

  t.Log( "#Score returns 0, false for no-win, non-full board" )
  board.Reset()
  score,final = minimax.Score( board, max )
  assert.Equal( t, score, 0 )
  assert.False( t, final )

  // with one level of recursion
  minimax.DepthLimit = 1

  t.Log( "#Score returns 1, true if 'max mark' can win in one move" )
  AddMarks( board, max, 0, 4 )
  score,final = minimax.Score( board, min )
  assert.Equal( t, score, 1 )
  assert.True( t, final )

  t.Log( "#Score returns -1, true if 'min mark' can win in one move" )
  AddMarks( board, min, 0, 4 )
  score,final = minimax.Score( board, max )
  assert.Equal( t, score, -1 )
  assert.True( t, final )
}

func TestMinimax_CurrentScore( t *testing.T ) {
  var minimax = NewMinimax()
  var board = NewBoard()
  var min, max = "O", "X"
  minimax.SetMinMaxMarks( min, max )

  t.Log( "#FinalScore returns 1, true for 'max' win" )
  AddMarks( board, max, 3, 4, 5 )
  score,final := minimax.FinalScore( board )
  assert.Equal( t, score, 1 )
  assert.True( t, final )

  t.Log( "#FinalScore returns -1, true for 'min' win" )
  AddMarks( board, min, 1, 4, 7 )
  score,final = minimax.FinalScore( board )
  assert.Equal( t, score, -1 )
  assert.True( t, final )

  t.Log( "#FinalScore returns 0, true if board is full without win" )
  AddMarks( board, max, 0, 1, 5, 6, 7 )
  AddMarks( board, min, 2, 3, 4, 8 )
  score,final = minimax.FinalScore( board )
  assert.Equal( t, score, 0 )
  assert.True( t, final )

  t.Log( "#FinalScore returns 0, false for no-win, non-full board" )
  board.Reset()
  score,final = minimax.FinalScore( board )
  assert.Equal( t, score, 0 )
  assert.False( t, final )
}

func TestMinimax_BestScoreForMark( t *testing.T ) {
  var minimax = NewMinimax()
  var min, max = "O", "X"
  minimax.SetMinMaxMarks( min, max )

  t.Log( "best score for min mark is -1" )
  assert.Equal( t, minimax.BestScoreForMark( min ), -1 )

  t.Log( "best score for max mark is 1" )
  assert.Equal( t, minimax.BestScoreForMark( max ), 1 )
}

func TestMinimax_BestOfScores( t *testing.T ) {
  var minimax = NewMinimax()
  var min, max = "O", "X"
  minimax.SetMinMaxMarks( min, max )

  allValues   := map[int]int{ 0:0, 1:1, 2:-1, 3:0 }
  highValues  := map[int]int{ 0:1, 1:0 }
  lowValues   := map[int]int{ 0:0, 2:-1 }

  t.Log( "best score for min mark is lowest" )
  assert.Equal( t, minimax.BestOfScores( allValues, min ), -1 )
  assert.Equal( t, minimax.BestOfScores( highValues, min ), 0 )
  assert.Equal( t, minimax.BestOfScores( lowValues, min ), -1 )

  t.Log( "best score for max mark is highest" )
  assert.Equal( t, minimax.BestOfScores( allValues, max ), 1 )
  assert.Equal( t, minimax.BestOfScores( highValues, max ), 1 )
  assert.Equal( t, minimax.BestOfScores( lowValues, max ), 0 )
}

func TestMinimax_SetMinMaxMarks( t *testing.T ) {
  var minimax = NewMinimax()

  t.Log( "sets marks to be used internally" )
  minimax.SetMinMaxMarks( "O", "X" )
  assert.Equal( t, minimax.BestScoreForMark( "O" ), -1 )
  assert.Equal( t, minimax.BestScoreForMark( "X" ), 1 )
  minimax.SetMinMaxMarks( "X", "O" )
  assert.Equal( t, minimax.BestScoreForMark( "O" ), 1 )
  assert.Equal( t, minimax.BestScoreForMark( "X" ), -1 )
}

func BenchmarkMinimax_ScoreAvailableMoves( b *testing.B ) {
  b.StopTimer()
  var minimax = NewMinimax()
  minimax.DepthLimit = 4
  var board = NewBoard()

  b.StartTimer()
  minimax.ScoreAvailableMoves( board, "X" )
}
