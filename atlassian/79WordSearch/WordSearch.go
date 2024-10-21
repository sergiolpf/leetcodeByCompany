package wordsearch

func Exist(board [][]byte, word string) bool {

	rows := len(board)

	if rows == 0 {
		return false
	}
	columns := len(board[0])
	wordSize := len(word)

	if rows == 1 && columns == 1 {
		return string(board[0][0]) == word
	}

	offSet := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var Backtracking func(row, col, index int) bool

	Backtracking = func(row, col, index int) bool {

		if index == wordSize {
			return true
		}

		if board[row][col] != byte(word[index]) {
			return false
		}

		temp := board[row][col]
		board[row][col] = '#'

		for _, offSetValues := range offSet {
			r, c := row+offSetValues[0], col+offSetValues[1]
			if (0 <= r && r < rows) && (0 <= c && c < columns) {
				if Backtracking(r, c, index+1) {
					return true
				}
			}
		}
		board[row][col] = temp

		return false
	}

	for r := range board {
		for c := range board[0] {
			if Backtracking(r, c, 0) {
				return true
			}
		}

	}
	return false
}
