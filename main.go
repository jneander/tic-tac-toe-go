package main

import "tictactoe/ttt"
import "os"

func main() {
  var ui = ttt.NewConsole( os.Stdin, os.Stdout )
  var game = ttt.NewGame()
  runner := ttt.ConsoleRunner{ game, ui }
  runner.Start()
}
