package ttt

import sassert "github.com/sdegutis/go.assert"
import "github.com/stretchrcom/testify/assert"
import "testing"

func TestFakeConsole( t *testing.T ) {
  var fake = NewFakeConsole()

  t.Log( "is substitutable for UI" )
  var ui UI = NewFakeConsole()
  assert.Equal( t, ui, ui )

  // Stubbing Method Responses

  t.Log( "stubs responses to #PromptMainMenu" )
  fake.StubPromptMainMenu( 3, 7, 4 )
  assert.Equal( t, fake.PromptMainMenu(), 3 )
  assert.Equal( t, fake.PromptMainMenu(), 7 )
  assert.Equal( t, fake.PromptMainMenu(), 4 )

  t.Log( "stubs responses to #PromptPlayerMove" )
  fake.StubPromptPlayerMove( 3, 7, 4 )
  assert.Equal( t, fake.PromptPlayerMove(), 3 )
  assert.Equal( t, fake.PromptPlayerMove(), 7 )
  assert.Equal( t, fake.PromptPlayerMove(), 4 )

  t.Log( "stubs repeat last response" )
  assert.Equal( t, fake.PromptMainMenu(), 4 )
  assert.Equal( t, fake.PromptPlayerMove(), 4 )

  // Spying on Methods

  t.Log( "#SpyOn spies on #PromptMainMenu" )
  assert.Equal( t, fake.TimesCalled( "PromptMainMenu" ), 0 )
  fake.SpyOn( "PromptMainMenu" )
  fake.StubPromptMainMenu( 1, 2, 3 )
  for i := 0; i < 3; i++ { fake.PromptMainMenu() }
  assert.Equal( t, fake.TimesCalled( "PromptMainMenu" ), 3 )

  t.Log( "#SpyOn spies on #PromptPlayerMove" )
  assert.Equal( t, fake.TimesCalled( "PromptPlayerMove" ), 0 )
  fake.SpyOn( "PromptPlayerMove" )
  fake.StubPromptPlayerMove( 1, 2, 3 )
  for i := 0; i < 3; i++ { fake.PromptPlayerMove() }
  assert.Equal( t, fake.TimesCalled( "PromptPlayerMove" ), 3 )

  t.Log( "#SpyOn can receive multiple methods to spy" )
  assert.Equal( t, fake.TimesCalled( "DisplayBoard" ), 0 )
  fake.SpyOn( "DisplayBoard", "DisplayAvailableSpaces" )

  t.Log( "#SpyOn spies on #DisplayBoard" )
  for i := 0; i < 2; i++ { fake.DisplayBoard( nil ) }
  assert.Equal( t, fake.TimesCalled( "DisplayBoard" ), 2 )

  t.Log( "#SpyOn spies on #DisplayAvailableSpaces" )
  for i := 0; i < 4; i++ { fake.DisplayAvailableSpaces( nil ) }
  assert.Equal( t, fake.TimesCalled( "DisplayAvailableSpaces" ), 4 )

  t.Log( "#SpyOn spies are logged independently" )
  assert.Equal( t, fake.TimesCalled( "PromptMainMenu" ), 3 )
  assert.Equal( t, fake.TimesCalled( "PromptPlayerMove" ), 3 )
  assert.Equal( t, fake.TimesCalled( "DisplayBoard" ), 2 )
  assert.Equal( t, fake.TimesCalled( "DisplayAvailableSpaces" ), 4 )

  // Resetting and Removing Spies

  t.Log( "#ResetSpies resets all spy counters" )
  fake.ResetSpies()
  assert.Equal( t, fake.TimesCalled( "PromptMainMenu" ), 0 )
  assert.Equal( t, fake.TimesCalled( "PromptPlayerMove" ), 0 )
  assert.Equal( t, fake.TimesCalled( "DisplayBoard" ), 0 )
  assert.Equal( t, fake.TimesCalled( "DisplayAvailableSpaces" ), 0 )

  t.Log( "#ResetSpies does not remove spies" )
  fake.DisplayBoard( nil )
  fake.DisplayAvailableSpaces( nil )
  assert.Equal( t, fake.TimesCalled( "DisplayBoard" ), 1 )
  assert.Equal( t, fake.TimesCalled( "DisplayAvailableSpaces" ), 1 )

  t.Log( "#RemoveSpies removes all spies" )
  fake.RemoveSpies()
  fake.DisplayBoard( nil )
  fake.DisplayAvailableSpaces( nil )
  assert.Equal( t, fake.TimesCalled( "DisplayBoard" ), 0 )
  assert.Equal( t, fake.TimesCalled( "DisplayAvailableSpaces" ), 0 )

  // Method Call Logging

  var logger []string
  fake = NewFakeConsole()

  t.Log( "initializes an internal logging array" )
  sassert.DeepEquals( t, *fake.SpyLog(), []string{} )

  t.Log( "#UseLog replaces the logging array" )
  fake.UseLogger( &logger )
  fake.SpyOn( "DisplayBoard" )
  for i := 0; i < 2; i++ { fake.DisplayBoard( nil ) }
  sassert.DeepEquals( t, logger, []string{ "DisplayBoard", "DisplayBoard" } )

  t.Log( "#SpyLog returns the logger being used" )
  assert.Equal( t, &logger, fake.SpyLog() )

  t.Log( "#ResetSpies clears any logging array" )
  fake.ResetSpies()
  assert.Equal( t, len( logger ), 0 )
  fake.DisplayBoard( nil )
  assert.Equal( t, len( logger ), 1 )

  t.Log( "#RemoveSpies replaces any logging array" )
  fake.RemoveSpies()
  fake.SpyOn( "DisplayBoard" )
  assert.Equal( t, len( logger ), 0 )
  fake.DisplayBoard( nil )
  assert.Equal( t, len( logger ), 0 )
  sassert.DeepEquals( t, *fake.SpyLog(), []string{ "DisplayBoard" } )
}
