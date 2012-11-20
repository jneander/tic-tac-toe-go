package tictactoe

import "strings"

type ConsoleUI struct {
  in  Reader
  out Writer
}

func ( c ConsoleUI ) PromptMainMenu() {
  c.out.WriteString( "Welcome to Tic Tac Toe in Go!\n" +
                     "Press any key to exit... " )
  for keys := ""; len(keys) == 0; keys = ReadInput( c.in ) {}
}

func ( c ConsoleUI ) DisplayBoard( b Board ) {
  spaces := b.Spaces()
  rows := make( []string, 3 )
  for i := range rows {
    rows[i] = strings.Join( spaces[i*3:i*3+3], "|" )
    rows[i] = strings.Replace( rows[i], b.Blank(), "_", -1 )
  }
  c.out.WriteString( strings.Join( rows, "\n" ) + "\n" )
}

func ReadInput( buffer Reader ) ( result string ) {
  var read = make( []byte, 128 )
  num, _ := buffer.Read( read )
  return string( read[:num] )
}
