package ttt

type consolePlayer struct {
  mark string
  console UI
}

func NewConsolePlayer( console UI ) Player {
  player := new( consolePlayer )
  player.console = console
  return player
}

func ( c consolePlayer ) Move( board Board ) int {
  available := board.SpacesWithMark( board.Blank() )
  return c.console.PromptPlayerMove( available... )
}

func ( c *consolePlayer ) SetMark( mark string ) {
  c.mark = mark
}

func ( c consolePlayer ) GetMark() string {
  return c.mark
}
