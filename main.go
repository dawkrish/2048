package main

import (
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board      [4][4]int
	emptyTiles [][2]int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "down":
			doesSomethingChange := false
			for i := 2; i >= 0; i-- {
				for j := 0; j < 4; j++ {
					idx1, idx2 := m.DownNonZeroIndex(i, j)
					if idx1 == -1 && idx2 == -1 {
						doesSomethingChange = true
						m.board[3][j] = m.board[i][j]
						m.board[i][j] = 0
					} else {
						if m.board[idx1][idx2] == m.board[i][j] {
							doesSomethingChange = true
							m.board[idx1][idx2] += m.board[i][j]
							m.board[i][j] = 0
						}
					}
				}
			}
			if doesSomethingChange {
				print("I occour")
				m.RecomputeEmptyTiles()

				numOfEmpTiles := len(m.emptyTiles)
				randTileIdx := m.emptyTiles[rand.Intn(numOfEmpTiles)]
				x := randTileIdx[0]
				y := randTileIdx[1]

				if rand.Float32() < 0.75 {
					m.board[x][y] = 2
				} else {
					m.board[x][y] = 4
				}

				m.emptyTiles = append(m.emptyTiles, [2]int{x, y})
			}
		case "up":
			doesSomethingChange := false
			for i := 1; i < 4; i++ {
				for j := 0; j < 4; j++ {
					idx1, idx2 := m.UpNonZeroIndex(i, j)
					if idx1 == -1 && idx2 == -1 {
						doesSomethingChange = true
						m.board[0][j] = m.board[i][j]
						m.board[i][j] = 0
					} else {
						if m.board[idx1][idx2] == m.board[i][j] {
							doesSomethingChange = true
							m.board[idx1][idx2] += m.board[i][j]
							m.board[i][j] = 0
						}
					}
				}
			}
			if doesSomethingChange {
				m.RecomputeEmptyTiles()

				numOfEmpTiles := len(m.emptyTiles)
				randTileIdx := m.emptyTiles[rand.Intn(numOfEmpTiles)]
				x := randTileIdx[0]
				y := randTileIdx[1]

				if rand.Float32() < 0.75 {
					m.board[x][y] = 2
				} else {
					m.board[x][y] = 4
				}

				m.emptyTiles = append(m.emptyTiles, [2]int{x, y})

			}
		case "left":
		case "right":

		}
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
	return m, nil
}

func (m model) View() string {
	s := "Welcome to 2048\n\n"
	s += "Use arrow keys to move the tiles.\nWhen two tiles having the same number touch, they join into one!\n\n"

	for i := range m.board {
		for j := range m.board[i] {
			val := m.board[i][j]
			if j == 0 {
				s += "|"
			}
			if val == 0 {
				s += fmt.Sprintf("    |")
			} else {
				s += fmt.Sprintf("%4d|", val)
			}
		}
		s += "\n"
	}

	s += "\nPress Ctrl+C or 'q' to exit\n"
	return s
}

func initialModel() model {
	board := [4][4]int{}
	x1, y1 := randomPosition()
	x2, y2 := randomPosition()

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

func randomPosition() (int, int) {
	return rand.Intn(4), rand.Intn(4)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
