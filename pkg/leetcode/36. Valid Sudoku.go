package leetcode

import "strconv"

func isValidSudoku(board [][]byte) bool {
	tem := make([]byte, 10)
	for i := 0; i < 9; i++ {
		tem = make([]byte, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if tem[byteToInt(board[i][j])] != 0 {
					return false
				} else {
					tem[byteToInt(board[i][j])] = 1
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		tem = make([]byte, 10)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				if tem[byteToInt(board[j][i])] != 0 {
					return false
				} else {
					tem[byteToInt(board[j][i])] = 1
				}
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tem = make([]byte, 10)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[k+3*i][l+3*j] != '.' {
						if tem[byteToInt(board[k+3*i][l+3*j])] != 0 {
							return false
						} else {
							tem[byteToInt(board[k+3*i][l+3*j])] = 1
						}
					}
				}
			}
		}
	}
	return true
}

func byteToInt(b byte) int {
	i, _ := strconv.Atoi(string(b))
	return i
}
