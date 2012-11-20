package tictactoe

type Board struct {
  spaces []string
}

func ( b Board ) Blank() string {
  return " "
}

func ( b *Board ) Spaces() []string {
  dup := make( []string, len( b.spaces ) )
  copy( dup, b.spaces )
  return dup
}

func ( b *Board ) Mark( pos int, mark string ) {
  if pos >= 0 && pos < len( b.spaces ) {
    b.spaces[ pos ] = mark
  }
}

func ( b *Board ) Reset() {
  setBoard( b )
}

func NewBoard() Board {
  b := new( Board )
  setBoard( b )
  return *b
}

func setBoard( b *Board ) {
  b.spaces = make( []string, 9 )
  for i := range b.spaces {
    b.spaces[ i ] = " "
  }
}
