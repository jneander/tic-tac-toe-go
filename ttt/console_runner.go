package ttt

type ConsoleRunner struct {
  Game Game
  UI UI
  Players []Player
}

func ( runner *ConsoleRunner ) Run() {
  if runner.UI.PromptMainMenu() != EXIT_GAME {
    runGameLoop( runner.Game, runner.Players, runner.UI )
  }
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
