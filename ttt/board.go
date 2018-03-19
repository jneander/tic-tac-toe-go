package ttt

// Board contains spaces
type Board struct {
	spaces []string
}

// Blank returns a whitespace
func (b Board) Blank() string {
	return " "
}

func (b *Board) Spaces() []string {
	dup := make([]string, len(b.spaces))
	copy(dup, b.spaces)
	return dup
}

func (b *Board) Mark(pos int, mark string) {
	if pos >= 0 && pos < len(b.spaces) {
		b.spaces[pos] = mark
	}
}

func (b Board) SpacesWithMark(mark string) []int {
	count, result := 0, make([]int, len(b.Spaces()))
	for i, v := range b.Spaces() {
		if v == mark {
			result[count] = i
			count++
		}
	}
	return result[:count]
}

// Reset clears the marks on the board
func (b *Board) Reset() {
	setBoard(b)
}

// NewBoard returns pointer to a new board
func NewBoard() *Board {
	b := new(Board)
	setBoard(b)
	return b
}

func setBoard(b *Board) {
	b.spaces = make([]string, 9)
	for i := range b.spaces {
		b.spaces[i] = " "
	}
}
