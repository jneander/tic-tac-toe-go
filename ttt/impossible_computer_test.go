package ttt

import (
	"testing"

	"github.com/stretchrcom/testify/assert"
)

func TestNewImpossibleComputer(t *testing.T) {
	var computer = NewImpossibleComputer()

	t.Log("is substitutable for Player")
	var player Player = computer
	assert.Equal(t, player, computer)

	t.Log("initializes with a Minimax instance")
	minimax, ok := computer.Minimax.(*Minimax)
	assert.True(t, ok)
	assert.NotEqual(t, minimax, (*Minimax)(nil))
}

func TestImpossibleComputer_Move(t *testing.T) {
	var computer = NewImpossibleComputer()
	var board = NewBoard()
	computer.SetMark("X")

	t.Log("#Move returns any winning move")
	AddMarks(board, "X", 1, 4)
	AddMarks(board, "O", 0, 2, 3)
	assert.Equal(t, computer.Move(*board), 7)

	t.Log("#Move blocks any winning move")
	board.Reset()
	AddMarks(board, "X", 0, 6)
	AddMarks(board, "O", 3, 4)
	assert.Equal(t, computer.Move(*board), 5)

	fakeMinimax := new(FakeMinimax)
	computer.Minimax = fakeMinimax

	t.Log("#Move uses the highest of Minimax scores")
	board.Reset()
	fakeMinimax.StubScores = map[int]int{1: 0, 3: 1, 4: -1, 5: 0}
	assert.Equal(t, computer.Move(*board), 3)
	fakeMinimax.StubScores = map[int]int{1: -1, 3: -1, 4: -1, 5: 0}
	assert.Equal(t, computer.Move(*board), 5)
}

func TestImpossibleComputer_Mark(t *testing.T) {
	var computer = NewImpossibleComputer()

	t.Log("implements GetMark and SetMark")
	computer.SetMark("X")
	assert.Equal(t, computer.GetMark(), "X")
	computer.SetMark("O")
	assert.Equal(t, computer.GetMark(), "O")

	fakeMinimax := new(FakeMinimax)
	computer.Minimax = fakeMinimax

	t.Log("sets marks on Minimax instance")
	computer.SetMark("X")
	assert.Equal(t, fakeMinimax.minMark, "O")
	assert.Equal(t, fakeMinimax.maxMark, "X")
	computer.SetMark("O")
	assert.Equal(t, fakeMinimax.minMark, "X")
	assert.Equal(t, fakeMinimax.maxMark, "O")
}

type FakeMinimax struct {
	StubScores map[int]int
	minMark    string
	maxMark    string
}

func (f *FakeMinimax) ScoreAvailableMoves(*Board, string) (map[int]int, bool) {
	return f.StubScores, false
}

func (f *FakeMinimax) SetMinMaxMarks(min string, max string) {
	f.minMark = min
	f.maxMark = max
}
