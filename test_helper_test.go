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

  t.Log( "IsEmpty returns true if Queue contains no elements" )
  assert.True( t, queue.IsEmpty(), "queue is not empty" )

  t.Log( "Length returns the number of elements in the Queue" )
  assert.Equal( t, queue.Length(), 0 )
  queue.Push( 6, 7, 8 )
  assert.Equal( t, queue.Length(), 3 )

  t.Log( "Queue works with multiple data types" )
  queue = new( Queue )
  queue.Push( "string", "next" )
  assert.Equal( t, queue.Pop(), "string" )
  assert.Equal( t, queue.Pop(), "next" )
}
