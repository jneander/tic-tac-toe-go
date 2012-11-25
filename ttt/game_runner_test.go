package ttt

import sassert "github.com/sdegutis/go.assert"
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
  SetInputs( &in, "9" )
  runner.Start()
  assert.Equal( t, game.Board().Spaces()[8], "X" )

  t.Log( "applies alternating marks for successive spaces" )
  game.Reset()
  SetInputs( &in, "3", "1", "4", "2", "5", "6", "8", "7", "9" )
  runner.Start()
  assert.Equal( t, game.Board().Spaces()[2], "X" )
  assert.Equal( t, game.Board().Spaces()[0], "O" )
  assert.Equal( t, game.Board().Spaces()[3], "X" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )

  t.Log( "stops applying moves when game is over" )
  game.Reset()
  MakeMoves( game, "X", 1, 4 )
  SetInputs( &in, "8", "9", "1" )
  runner.Start()
  assert.True( t, game.IsOver() )
  assert.Equal( t, game.Board().Spaces()[8], game.Board().Blank() )
  assert.Equal( t, game.Board().Spaces()[0], game.Board().Blank() )

  t.Log( "rejects invalid moves" )
  game.Reset()
  SetInputs( &in, "1", "1", "2", "4", "5", "4", "7" )
  runner.Start()
  assert.Equal( t, game.Board().Spaces()[0], "X" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )

  t.Log( "displays the board before prompting for moves" )
  game.Reset()
  mui := NewConsoleUISpy( &in, &out )
  mui.SpyOn( "DisplayAvailableSpaces", "PromptPlayerMove", "DisplayBoard" )
  runner = prepareRunner( mui, game )
  SetInputs( &in, "1", "2", "4", "5", "7" )
  runner.Start()
  expected := []string{ "DisplayAvailableSpaces", "PromptPlayerMove",
                        "DisplayAvailableSpaces", "PromptPlayerMove" }
  sassert.DeepEquals( t, mui.methodCalls[:4], expected )

  t.Log( "displays the board after the game is over" )
  assert.Equal( t, mui.methodCalls[len(mui.methodCalls) - 1], "DisplayBoard" )
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

func NewConsoleUISpy( in Reader, out Writer ) *consoleSpy {
  spy := new( consoleSpy )
  spy.ui = new( ConsoleUI )
  spy.ui.in = in
  spy.ui.out = out
  spy.activeSpies = make( map[string]bool )
  return spy
}

type consoleSpy struct {
  ui *ConsoleUI
  methodCalls []string
  activeSpies map[string]bool
}

func ( spy *consoleSpy ) LogMethodCall( call string ) {
  newLog := make( []string, len( spy.methodCalls ) + 1 )
  copy( newLog, spy.methodCalls )
  newLog[ len(spy.methodCalls) ] = call
  spy.methodCalls = newLog
}

func ( spy *consoleSpy ) SpyOn( methods ...string ) {
  for _,v := range methods {
    spy.activeSpies[ v ] = true
  }
}

func ( spy *consoleSpy ) DisplayAvailableSpaces( board *Board ) {
  if spy.activeSpies[ "DisplayAvailableSpaces" ] {
    spy.LogMethodCall( "DisplayAvailableSpaces" )
  }
}
func ( spy *consoleSpy ) DisplayBoard( board *Board ) {
  if spy.activeSpies[ "DisplayBoard" ] {
    spy.LogMethodCall( "DisplayBoard" )
  }
}
func ( spy *consoleSpy ) PromptMainMenu() int {
  if spy.activeSpies[ "PromptMainMenu" ] {
    spy.LogMethodCall( "PromptMainMenu" )
  }
  return spy.ui.PromptMainMenu()
}

func ( spy *consoleSpy ) PromptPlayerMove ( valid ...interface{} ) int {
  if spy.activeSpies[ "PromptPlayerMove" ] {
    spy.LogMethodCall( "PromptPlayerMove" )
  }
  return spy.ui.PromptPlayerMove( valid... )
}
