package tictactoe

import "github.com/stretchrcom/testify/assert"
import "testing"

func TestQueue( t *testing.T ) {
  queue := new( Queue )

  t.Log( "Queue can push single values in" )
  queue.Push( 1 )
  assert.Equal( t, queue.Pop(), 1 )

  t.Log( "Queue can push multiple values in" )
  queue.Push( 3, 4 )
  assert.Equal( t, queue.Pop(), 3 )
  assert.Equal( t, queue.Pop(), 4 )

  t.Log( "Queue panics when empty and Popped" )
  assert.Panics( t, func(){ queue.Pop() } )

  t.Log( "Queue works with different data types" )
  queue = new( Queue )
  queue.Push( "string", "next" )
  assert.Equal( t, queue.Pop(), "string" )
  assert.Equal( t, queue.Pop(), "next" )
}
