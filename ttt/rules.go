package ttt

type Rules struct{}

func NewRules() *Rules {
	return new(Rules)
}

func (r *Rules) MarkHasWinningSolution(board *Board, mark string) bool {
	for _, set := range winningSets() {
		marks := marksForSet(board, set)
		if allMarksMatch(marks, mark) {
			return true
		}
	}
	return false
}

// PRIVATE

func marksForSet(board *Board, set []int) []string {
	result := make([]string, len(set))
	for i, v := range set {
		result[i] = board.Spaces()[v]
	}
	return result
}

func allMarksMatch(marks []string, mark string) bool {
	result := true
	for _, v := range marks {
		result = result && v == mark
	}
	return result
}

func winningSets() [][]int {
	return [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6, 7, 8},
		[]int{0, 3, 6}, []int{1, 4, 7}, []int{2, 5, 8},
		[]int{0, 4, 8}, []int{2, 4, 6}}
}
