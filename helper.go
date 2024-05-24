package main

import "math/rand"

func randomPosition() (int, int) {
	return rand.Intn(4), rand.Intn(4)
}

func (m model) RecomputeEmptyTiles() [][2]int {
	emptyTiles := [][2]int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m.board[i][j] == 0 {
				emptyTiles = append(emptyTiles, [2]int{i, j})
			}
		}
	}
	return emptyTiles
}

func (m model) DownNonZeroIndex(currentRow, currentCol int) int {
	for row := currentRow + 1; row < 4; row++ {
		if m.board[row][currentCol] != 0 {
			return row
		}
	}
	// The below column is empty
	return -1
}

func (m model) UpNonZeroIndex(currentRow, currentCol int) int {
	for row := currentRow - 1; row >= 0; row-- {
		if m.board[row][currentCol] != 0 {
			return row
		}
	}
	// The above column is empty
	return -1
}

func (m model) LeftNonZeroIndex(currentRow, currentCol int) int {
	for col := currentCol - 1; col >= 0; col-- {
		if m.board[currentRow][col] != 0 {
			return col
		}
	}

	return -1
}

func (m model) RightNonZeroIndex(currentRow, currentCol int) int {
	for col := currentCol + 1; col < 4; col++ {
		if m.board[currentRow][col] != 0 {
			return col
		}
	}

	return -1
}
