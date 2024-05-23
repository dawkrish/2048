package main

func (m model) RecomputeEmptyTiles() {
	emptyTiles := [][2]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m.board[i][j] == 0 {
				emptyTiles = append(emptyTiles, [2]int{i, j})
			}
		}
	}
	m.emptyTiles = emptyTiles
}

func (m model) DownNonZeroIndex(currentRow, currentCol int) (int, int) {
	for row := currentRow + 1; row < 4; row++ {
		if m.board[row][currentCol] != 0 {
			return row, currentCol
		}
	}
	// The below column is empty
	return -1, -1
}

func (m model) UpNonZeroIndex(currentRow, currentCol int) (int, int) {
	for row := currentRow - 1; row >= 0; row-- {
		if m.board[row][currentCol] != 0 {
			return row, currentCol
		}
	}
	// The above column is empty
	return -1, -1
}
