package ttt

type GameRunner struct {
  Game Game
  UI UI
}

func ( runner *GameRunner ) Start() {
  p1, p2 := "X", "O"
  for !runner.Game.IsOver() {
    move := runner.UI.PromptPlayerMove()
    if runner.Game.IsValidMove( move ) {
      runner.Game.ApplyMove( move, p1 )
      p1, p2 = p2, p1
    }
  }
}
