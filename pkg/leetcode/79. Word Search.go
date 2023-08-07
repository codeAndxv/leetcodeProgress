package leetcode

func exist(board [][]byte, word string) bool {
	record := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		record[i] = make([]bool, len(board[0]))
	}
	result := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			result = result || existBase(board, word, i, j, 0, record)
			if result {
				return true
			}
		}
	}
	return result
}

func existBase(board [][]byte, word string, x int, y int, currentIndex int, record [][]bool) bool {
	if record[x][y] || board[x][y] != word[currentIndex] {
		return false
	}
	if !record[x][y] && board[x][y] == word[currentIndex] && currentIndex == len(word)-1 {
		return true
	}
	record[x][y] = true
	result := false
	if x > 0 {
		result = result || existBase(board, word, x-1, y, currentIndex+1, record)
	}
	if x < len(board)-1 {
		result = result || existBase(board, word, x+1, y, currentIndex+1, record)
	}
	if y > 0 {
		result = result || existBase(board, word, x, y-1, currentIndex+1, record)
	}
	if y < len(board[0])-1 {
		result = result || existBase(board, word, x, y+1, currentIndex+1, record)
	}
	record[x][y] = false
	return result
}
