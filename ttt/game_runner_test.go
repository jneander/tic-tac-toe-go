package ttt

import "github.com/stretchrcom/testify/assert"
import "testing"
import "bytes"

func TestGameRunnerStart( t *testing.T ) {
  var in  bytes.Buffer
  var out bytes.Buffer
  var ui = &ConsoleUI{ &in, &out }
  var game = NewGame()
  runner := prepareRunner( ui, game )

  t.Log( "GameRunner stores references to UI and Game" )
  assert.Equal( t, runner.UI, ui )
  assert.Equal( t, runner.Game, game )

  t.Log( "applies a mark to the selected space" )
  MakeMoves( game, "X", 0, 1, 5, 6 )
  MakeMoves( game, "O", 2, 3, 4, 7 )
  SetInputs( &in, "8" )
  runner.Start()
  assert.Equal( t, game.Board().Spaces()[8], "X" )

  t.Log( "applies alternating marks for successive spaces" )
  game.Reset()
  SetInputs( &in, "2", "0", "3", "1", "4", "5", "7", "6", "8" )
  runner.Start()
  assert.Equal( t, game.Board().Spaces()[2], "X" )
  assert.Equal( t, game.Board().Spaces()[0], "O" )
  assert.Equal( t, game.Board().Spaces()[3], "X" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )

  t.Log( "stops applying moves when game is over" )
  game.Reset()
  MakeMoves( game, "X", 1, 4 )
  SetInputs( &in, "7", "8", "0" )
  runner.Start()
  assert.True( t, game.IsOver() )
  assert.Equal( t, game.Board().Spaces()[8], game.Board().Blank() )
  assert.Equal( t, game.Board().Spaces()[0], game.Board().Blank() )

  t.Log( "rejects invalid moves" )
  game.Reset()
  SetInputs( &in, "0", "0", "1", "3", "4", "3", "6" )
  runner.Start()
  assert.Equal( t, game.Board().Spaces()[0], "X" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )
}

func prepareRunner( ui UI, game Game ) *GameRunner {
  runner := new( GameRunner )
  runner.UI = ui
  runner.Game = game
  return runner
}

func MakeMoves( game Game, mark string, spaces ...int ) {
  for _,i := range spaces {
    game.ApplyMove( i, mark )
  }
}

func SetInputs( input *bytes.Buffer, data ...string ) {
  input.Reset();
  var result string
  for _,v := range data {
    result += v + "\n"
  }
  input.WriteString( result )
}
