package ttt

const (
  EXIT_GAME = iota
  PLAYER_FIRST
  COMPUTER_FIRST
)

type UI interface {
  PromptMainMenu() int
  DisplayBoard( *Board )
  PromptPlayerMove( ...int ) int
  DisplayAvailableSpaces( *Board )
  DisplayGameOver( Game )
}

type Reader interface {
  Read(b []byte) (n int, err error)
}

type Writer interface {
  WriteString(s string) (ret int, err error)
}
