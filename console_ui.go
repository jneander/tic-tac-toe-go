package tictactoe


type ConsoleUI struct {
  in  Reader
  out Writer
}

func ( c ConsoleUI ) PromptMainMenu() {
  c.out.WriteString( "Welcome to Tic Tac Toe in Go!\n" +
                     "Press any key to exit... " )
  for keys := ""; len(keys) == 0; keys = ReadInput( c.in ) {}
}

func ReadInput( buffer Reader ) ( result string ) {
  var read = make( []byte, 128 )
  num, _ := buffer.Read( read )
  return string( read[:num] )
}
