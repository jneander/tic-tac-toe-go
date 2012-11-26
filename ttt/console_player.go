package ttt

type ConsolePlayer struct {
  mark string
  console UI
}

func NewConsolePlayer( console UI ) *ConsolePlayer {
  player := new( ConsolePlayer )
  player.console = console
  return player
}

func ( c ConsolePlayer ) Move( board Board ) int {
  available := board.SpacesWithMark( board.Blank() )
  return c.console.PromptPlayerMove( available... )
}

func ( c *ConsolePlayer ) SetMark( mark string ) {
  c.mark = mark
}

func ( c ConsolePlayer ) GetMark() string {
  return c.mark
}
