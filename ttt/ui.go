package ttt

const (
  EXIT_GAME = iota
)

type UI interface {
  PromptMainMenu() int
  DisplayBoard( *Board )
  PromptPlayerMove( ...interface{} ) int
  DisplayAvailableSpaces( *Board )
}

type Reader interface {
  Read(b []byte) (n int, err error)
}

type Writer interface {
  WriteString(s string) (ret int, err error)
}
