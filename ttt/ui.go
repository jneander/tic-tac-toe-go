package ttt

type UI interface {
  PromptMainMenu()
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
