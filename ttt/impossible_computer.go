package ttt

type ImpossibleComputer struct {
	mark    string
	Minimax BoardScorer
}

type BoardScorer interface {
	ScoreAvailableMoves(*Board, string) (map[int]int, bool)
	SetMinMaxMarks(string, string)
}

func NewImpossibleComputer() *ImpossibleComputer {
	computer := new(ImpossibleComputer)
	computer.Minimax = NewMinimax()
	return computer
}

func (c *ImpossibleComputer) Move(board Board) int {
	moveScores, _ := c.Minimax.ScoreAvailableMoves(&board, c.mark)
	bestMove, bestScore := -1, -1
	for move, score := range moveScores {
		if score > bestScore {
			bestMove, bestScore = move, score
		}
	}
	return bestMove
}

func (c *ImpossibleComputer) SetMark(mark string) {
	c.mark = mark
	marks := map[string]string{"X": "O", "O": "X"}
	c.Minimax.SetMinMaxMarks(marks[mark], mark)
}

func (c *ImpossibleComputer) GetMark() string {
	return c.mark
}
