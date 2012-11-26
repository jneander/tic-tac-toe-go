package ttt

import "github.com/stretchrcom/testify/assert"
import "testing"

func TestFakePlayer( t *testing.T ) {
  var fake = new( FakePlayer )
  var board = NewBoard()

  t.Log( "#StubMoves stubs responses to #Move" )
  fake.StubMoves( 3, 7, 4 )
  assert.Equal( t, fake.Move( *board ), 3 )
  assert.Equal( t, fake.Move( *board ), 7 )
  assert.Equal( t, fake.Move( *board ), 4 )

  t.Log( "implements GetMark and SetMark" )
  fake.SetMark( "X" )
  assert.Equal( t, fake.GetMark(), "X" )
  fake.SetMark( "O" )
  assert.Equal( t, fake.GetMark(), "O" )
}
