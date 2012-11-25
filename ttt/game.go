package ttt

type Game interface {
  Board() *Board
  IsOver() bool
  IsValidMove( int ) bool
  ApplyMove( int, string )
}

type game struct {
  board *Board
}

func NewGame() *game {
  g := new( game )
  g.board = NewBoard()
  return g
}

func ( g *game ) Board() *Board {
  return g.board
}

func ( g *game ) IsOver() bool {
  return winningSetExists( g.board ) || boardIsFull( g.board )
}

func ( g *game ) IsValidMove( space int ) bool {
  board := g.board
  return board.Spaces()[ space ] == board.Blank()
}

func ( g *game ) ApplyMove( pos int, mark string ) {
  if ( g.IsValidMove( pos ) ) {
    g.board.Mark( pos, mark )
  }
}

func ( g *game ) Reset() {
  g.board.Reset()
}

// PRIVATE

func boardIsFull( board *Board ) bool {
  for _,mark := range board.Spaces() {
    if mark == board.Blank() { return false }
  }
  return true
}

func winningSetExists( board *Board ) bool {
  exists := false
  for _,set := range solutions() {
    exists = exists || allSpacesMatch( board, set )
  }
  return exists
}

func allSpacesMatch( board *Board, pos []int ) bool {
  spaces := board.Spaces()
  mark := spaces[ pos[ 0 ] ]
  result := mark != board.Blank()
  for _,i := range pos {
    result = result && spaces[ i ] == mark
  }
  return result
}

func solutions() [][]int {
  return [][]int{ []int{ 0, 1, 2 }, []int{ 3, 4, 5 }, []int{ 6, 7, 8 },
                  []int{ 0, 3, 6 }, []int{ 1, 4, 7 }, []int{ 2, 5, 8 },
                  []int{ 0, 4, 8 }, []int{ 2, 4, 6 } }
}
