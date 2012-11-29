package ttt

func NewFakeConsole() *fakeConsole {
  console := new( fakeConsole )
  console.RemoveSpies()
  return console
}

type fakeConsole struct {
  altLogger *[]string
  spyCalls map[string]int
  activeSpies []string
  stubbedMainMenu []int
  stubbedPlayerMoves []int
}

func (f *fakeConsole) StubPromptMainMenu( responses ...int ) {
  f.stubbedMainMenu = responses
}

func (f *fakeConsole) StubPromptPlayerMove( moves ...int ) {
  f.stubbedPlayerMoves = moves
}

func (f *fakeConsole) PromptMainMenu() int {
  f.logSpyCall( "PromptMainMenu" )
  result := f.stubbedMainMenu[0]
  if len( f.stubbedMainMenu ) > 1 {
    f.stubbedMainMenu = f.stubbedMainMenu[1:]
  }
  return result
}

func (f *fakeConsole) PromptPlayerMove( valid ...int ) int {
  f.logSpyCall( "PromptPlayerMove" )
  result := f.stubbedPlayerMoves[0]
  if len( f.stubbedPlayerMoves ) > 1 {
    f.stubbedPlayerMoves = f.stubbedPlayerMoves[1:]
  }
  return result
}

func (f *fakeConsole) DisplayAvailableSpaces( board *Board ) {
  f.logSpyCall( "DisplayAvailableSpaces" )
}

func (f *fakeConsole) DisplayBoard( board *Board ) {
  f.logSpyCall( "DisplayBoard" )
}

func (f *fakeConsole) DisplayGameOver( game Game ) {
  f.logSpyCall( "DisplayGameOver" )
}

func (f *fakeConsole) UseLogger( logger *[]string ) {
  f.altLogger = logger
}

func (f *fakeConsole) SpyLog() *[]string {
  return f.altLogger
}

func (f *fakeConsole) SpyOn( method ...string ) {
  f.activeSpies = append( f.activeSpies, method... )
}

func (f *fakeConsole) TimesCalled( method string ) int {
  if _,ok := f.spyCalls[ method ]; ok {
    return f.spyCalls[ method ]
  }
  return 0
}

func (f *fakeConsole) ResetSpies() {
  f.spyCalls = make( map[string]int )
  if f.altLogger != nil { *f.altLogger = []string{} }
}

func (f *fakeConsole) RemoveSpies() {
  f.ResetSpies()
  f.activeSpies = make( []string, 0 )
  f.altLogger = &([]string{})
}

func (f *fakeConsole) logSpyCall( method string ) {
  for _,v := range f.activeSpies {
    if v == method {
      f.spyCalls[ method ]++
      if f.altLogger != nil {
        *f.altLogger = append( (*f.altLogger)[:], method )
      }
      break
    }
  }
}
