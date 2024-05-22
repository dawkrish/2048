package main

func (m model) DownNonZeroIndex(currentRow, currentCol int) (int, int) {
	for row := currentRow; row < 4; row++ {
		if m.board[row][currentCol] != 0 {
			return row, currentCol
		}
	}
	return -1, -1
}
