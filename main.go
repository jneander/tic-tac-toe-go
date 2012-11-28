package main

import "tictactoe/ttt"
import "os"

func main() {
  var ui = ttt.NewConsole( os.Stdin, os.Stdout )
  var game = ttt.NewGame()
  var p1 = ttt.NewConsolePlayer( ui )
  var p2 = ttt.NewImpossibleComputer()
  var players = []ttt.Player{ p1, p2 }
  p1.SetMark( "O" )
  p2.SetMark( "X" )

  runner := ttt.ConsoleRunner{ game, ui, players }
  runner.Run()
}
