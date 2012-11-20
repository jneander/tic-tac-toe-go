package tictactoe

import "bytes"

type Console struct {
  in  *bytes.Buffer
  out *bytes.Buffer
}

func ( c Console ) PromptMainMenu() {
  c.out.WriteString( "Welcome to Tic Tac Toe in Go!\n" +
                     "Press any key to exit... " )
  for keys := ""; len(keys) == 0; keys = ReadBuffer( c.in ) {}
}
