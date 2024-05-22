package main

import (
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board [4][4]int
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
			for i := 2; i >= 0; i-- {
				for j := 0; j < 4; j++ {

					if m.board[i+1][j] == 0 {
						m.board[i+1][j] = m.board[i][j]
						m.board[i][j] = 0
					} else {
						if m.board[i+1][j] == m.board[i][j] {
							m.board[i+1][j] += m.board[i][j]
							m.board[i][j] = 0
						}
					}

				}
			}
		case "up":
		case "left":
		case "right":
		}

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

	return model{
		board: board,
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
