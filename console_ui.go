package tictactoe


type Console struct {
  in  Reader
  out Writer
}

func ( c Console ) PromptMainMenu() {
  c.out.WriteString( "Welcome to Tic Tac Toe in Go!\n" +
                     "Press any key to exit... " )
  for keys := ""; len(keys) == 0; keys = ReadInput( c.in ) {}
}
