package ttt

type ConsoleRunner struct {
  Game Game
  UI UI
  Players []Player
}

func ( runner *ConsoleRunner ) Run() {
  selection := runner.UI.PromptMainMenu()
  if selection == EXIT_GAME { return }
  if selection == COMPUTER_FIRST {
    runner.Players[0], runner.Players[1] = runner.Players[1], runner.Players[0]
  }
  runGameLoop( runner.Game, runner.Players, runner.UI )
}

func runGameLoop( game Game, players []Player, console UI ) {
  current, next := players[0], players[1]
  for !game.IsOver() {
    console.DisplayAvailableSpaces( game.Board() )
    move := current.Move( *game.Board() )
    if game.IsValidMove( move ) {
      game.ApplyMove( move, current.GetMark() )
      current, next = next, current
    }
  }
  console.DisplayBoard( game.Board() )
}
