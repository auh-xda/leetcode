func isValidSudoku(board [][]byte) bool {
    cellRange := func(row int, col int) (int, int, int, int) {
		rStart := (row / 3) * 3
		rEnd := (((row / 3) + 1) * 3) - 1

		cStart := (col / 3) * 3
		cEnd := (((col / 3) + 1) * 3) - 1

		return rStart, rEnd, cStart, cEnd
	}

	existsInRow := func(board [][]byte, row int, pos int, val byte) bool {
		for pos < 8 {
			if board[row][pos+1] == val {
				return true
			}
			pos++
		}

		return false
	}

	existsInColumn := func(board [][]byte, pos int, col int, val byte) bool {
		for pos < 8 {
			if board[pos+1][col] == val {
				return true
			}
			pos++
		}

		return false
	}

	existsInCell := func(board [][]byte, row int, col int, val byte) bool {
		xStart, xEnd, yStart, yEnd := cellRange(row, col)

		for x := xStart; x <= xEnd; x++ {
			for y := yStart; y <= yEnd; y++ {
				if x == row && y == col {
					continue
				}

				if board[x][y] == val {
					return true
				}
			}
		}

		return false
	}

	for i := range 9 {
		for j := range 9 {
			current := board[i][j]
			if current != '.' {
				if existsInRow(board, i, j, current) || existsInColumn(board, i, j, current) || existsInCell(board, i, j, current) {
					fmt.Printf("Failing for %d, %d", i, j)
					return false
				}
			}
		}
	}

	return true
}