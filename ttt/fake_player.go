package ttt

type FakePlayer struct {
	mark         string
	stubbedMoves []int
}

func (f *FakePlayer) StubMoves(moves ...int) {
	f.stubbedMoves = moves
}

func (f *FakePlayer) Move(board Board) int {
	result := f.stubbedMoves[0]
	f.stubbedMoves = f.stubbedMoves[1:]
	return result
}

func (f *FakePlayer) GetMark() string {
	return f.mark
}

func (f *FakePlayer) SetMark(mark string) {
	f.mark = mark
}
