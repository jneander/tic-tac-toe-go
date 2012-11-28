package ttt

import "github.com/stretchrcom/testify/assert"
import "testing"

func TestConsoleRunnerRun( t *testing.T ) {
  var console = NewFakeConsole()
  var game = NewGame()
  var p1, p2 = new( FakePlayer ), new( FakePlayer )
  var players = []Player{ p1, p2 }
  p1.SetMark( "X" )
  p2.SetMark( "O" )

  var runner = ConsoleRunner{ game, console, players }

  // Exit the game
  console.StubPromptMainMenu( EXIT_GAME )

  t.Log( "exits immediately upon 'exit game' menu selection" )
  console.SpyOn( "DisplayAvailableSpaces", "DisplayBoard" )
  runner.Run()
  assert.Equal( t, len( *console.SpyLog() ), 0 )

  // Enter a 'Player goes first' loop
  console.StubPromptMainMenu( PLAYER_FIRST )

  t.Log( "applies alternating marks for successive spaces" )
  game.Reset()
  p1.StubMoves( 2, 3, 4, 7, 8 )
  p2.StubMoves( 0, 1, 5, 6 )
  runner.Run()
  assert.Equal( t, game.Board().Spaces()[2], "X" )
  assert.Equal( t, game.Board().Spaces()[0], "O" )
  assert.Equal( t, game.Board().Spaces()[3], "X" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )

  t.Log( "stops applying moves when game is over" )
  game.Reset()
  game.ApplyMove( 1, "X" )
  game.ApplyMove( 4, "X" )
  p1.StubMoves( 7, 0 )
  p2.StubMoves( 8 )
  runner.Run()
  assert.True( t, game.IsOver() )
  assert.Equal( t, game.Board().Spaces()[8], game.Board().Blank() )
  assert.Equal( t, game.Board().Spaces()[0], game.Board().Blank() )

  t.Log( "rejects invalid moves" )
  game.Reset()
  p1.StubMoves( 0, 1, -1, 4, 8 )
  p2.StubMoves( 0, 1, 6 )
  runner.Run()
  assert.Equal( t, game.Board().Spaces()[0], "X" )
  assert.Equal( t, game.Board().Spaces()[1], "O" )
  assert.Equal( t, game.Board().Spaces()[4], "X" )

  t.Log( "displays the available spaces before each move" )
  game.Reset()
  console.ResetSpies()
  console.SpyOn( "DisplayAvailableSpaces" )
  p1.StubMoves( 0, 1, -1, 4, 8 )
  p2.StubMoves( 0, 1, 6 )
  runner.Run()
  // TODO confirm that this call takes place before each move request
  assert.Equal( t, console.TimesCalled( "DisplayAvailableSpaces" ), 8 )

  t.Log( "displays the board when the game is over" )
  console.SpyOn( "DisplayBoard" )
  runner.Run()
  log := *console.SpyLog()
  assert.Equal( t, log[ len( log ) - 1 ], "DisplayBoard" )
}
