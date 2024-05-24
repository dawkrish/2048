package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board      [4][4]int
	emptyTiles [][2]int
}

func initialModel() model {
	board := [4][4]int{}
	x1, y1 := randomPosition()
	x2, y2 := randomPosition()
	// x1, y1 := 2, 1
	// x2, y2 := 2, 0

	board[x1][y1] = 2
	board[x2][y2] = 2

	emptyTiles := [][2]int{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if board[i][j] == 0 {
				emptyTiles = append(emptyTiles, [2]int{i, j})
			}
		}
	}

	return model{
		board:      board,
		emptyTiles: emptyTiles,
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
