package main

import "tic-tac-toe-go/ttt"
import "os"

func main() {
  var ui = ttt.NewConsole( os.Stdin, os.Stdout )
  var game = ttt.NewGame()
  var p1 = ttt.NewConsolePlayer( ui )
  var p2 = ttt.NewImpossibleComputer()
  var players = []ttt.Player{ p1, p2 }
  p1.SetMark( "X" )
  p2.SetMark( "O" )

  runner := ttt.ConsoleRunner{ game, ui, players }
  runner.Run()
}
