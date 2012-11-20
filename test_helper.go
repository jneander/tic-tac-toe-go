package tictactoe

import "bytes"

func ReadBuffer( buffer *bytes.Buffer ) ( result string ) {
  var read = make( []byte, 128 )
  num, _ := buffer.Read( read )
  return string( read[:num] )
}
