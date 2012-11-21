package tictactoe

type Game struct {
  Board Board
}

func ( g Game ) IsOver() bool {
  return winningSetExists( g.Board ) || boardIsFull( g.Board )
}

func NewGame() Game {
  g := new( Game )
  return *g
}

func boardIsFull( board Board ) bool {
  for _,mark := range board.Spaces() {
    if mark == board.Blank() { return false }
  }
  return true
}

func winningSetExists( board Board ) bool {
  exists := false
  for _,set := range solutions() {
    exists = exists || allSpacesMatch( board, set )
  }
  return exists
}

func allSpacesMatch( board Board, pos []int ) bool {
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
