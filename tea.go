package main

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
)

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
					if m.board[i][j] != 0 {
						var newRowIdx int
						var newVal int = m.board[i][j]
						idx := m.DownNonZeroIndex(i, j)

						if idx == -1 {
							newRowIdx = 3
						} else {
							if m.board[idx][j] == m.board[i][j] {
								newRowIdx = idx
								newVal *= 2
							} else {
								newRowIdx = idx - 1
							}
						}

						// fmt.Print("NewRowIdx: ", newRowIdx)
						// fmt.Print("NewVal: ", newVal)

						if newRowIdx != i {
							m.board[newRowIdx][j] = newVal
							m.board[i][j] = 0
							doesSomethingChange = true
						}
					}
				}
			}
			if doesSomethingChange {
				m.emptyTiles = m.RecomputeEmptyTiles()

				numOfEmpTiles := len(m.emptyTiles)
				randTileIdx := m.emptyTiles[rand.Intn(numOfEmpTiles)]
				x := randTileIdx[0]
				y := randTileIdx[1]

				if rand.Float32() < 0.85 {
					m.board[x][y] = 2
				} else {
					m.board[x][y] = 4
				}

				m.emptyTiles = m.RecomputeEmptyTiles()
			}
		case "up":
			doesSomethingChange := false
			for i := 1; i < 4; i++ {
				for j := 0; j < 4; j++ {
					if m.board[i][j] != 0 {
						idx := m.UpNonZeroIndex(i, j)
						var newRowIdx int
						var newVal int = m.board[i][j]
						if idx == -1 {
							newRowIdx = 0
						} else {
							if m.board[idx][j] == m.board[i][j] {
								newRowIdx = idx
								newVal *= 2
							} else {
								newRowIdx = idx + 1
							}
						}
						if newRowIdx != i {
							m.board[newRowIdx][j] = newVal
							m.board[i][j] = 0
							doesSomethingChange = true
						}

					}
				}
			}
			if doesSomethingChange {
				m.emptyTiles = m.RecomputeEmptyTiles()

				numOfEmpTiles := len(m.emptyTiles)
				randTileIdx := m.emptyTiles[rand.Intn(numOfEmpTiles)]
				x := randTileIdx[0]
				y := randTileIdx[1]

				if rand.Float32() < 0.85 {
					m.board[x][y] = 2
				} else {
					m.board[x][y] = 4
				}

				m.emptyTiles = m.RecomputeEmptyTiles()
			}
		case "left":
			doesSomethingChange := false

			for i := 1; i < 4; i++ {
				for j := 0; j < 4; j++ {
					if m.board[j][i] != 0 {
						idx := m.LeftNonZeroIndex(j, i)
						var newColIdx int
						var newVal int = m.board[j][i]

						if idx == -1 {
							newColIdx = 0
						} else {
							if m.board[j][idx] == m.board[j][i] {
								newColIdx = idx
								newVal *= 2
							} else {
								newColIdx = idx + 1
							}
						}

						if newColIdx != i {
							m.board[j][newColIdx] = newVal
							m.board[j][i] = 0
							doesSomethingChange = true
						}
					}
				}
			}

			if doesSomethingChange {
				m.emptyTiles = m.RecomputeEmptyTiles()

				numOfEmpTiles := len(m.emptyTiles)
				randTileIdx := m.emptyTiles[rand.Intn(numOfEmpTiles)]
				x := randTileIdx[0]
				y := randTileIdx[1]

				if rand.Float32() < 0.85 {
					m.board[x][y] = 2
				} else {
					m.board[x][y] = 4
				}

				m.emptyTiles = m.RecomputeEmptyTiles()
			}
		case "right":
			doesSomethingChange := false

			for i := 2; i >= 0; i-- {
				for j := 0; j < 4; j++ {
					if m.board[j][i] != 0 {
						idx := m.RightNonZeroIndex(j, i)
						var newColIdx int
						var newVal int = m.board[j][i]

						if idx == -1 {
							newColIdx = 3
						} else {
							if m.board[j][idx] == m.board[j][i] {
								newColIdx = idx
								newVal *= 2
							} else {
								newColIdx = idx - 1
							}
						}

						if newColIdx != i {
							m.board[j][newColIdx] = newVal
							m.board[j][i] = 0
							doesSomethingChange = true
						}
					}
				}
			}

			if doesSomethingChange {
				m.emptyTiles = m.RecomputeEmptyTiles()

				numOfEmpTiles := len(m.emptyTiles)
				randTileIdx := m.emptyTiles[rand.Intn(numOfEmpTiles)]
				x := randTileIdx[0]
				y := randTileIdx[1]

				if rand.Float32() < 0.85 {
					m.board[x][y] = 2
				} else {
					m.board[x][y] = 4
				}

				m.emptyTiles = m.RecomputeEmptyTiles()
			}
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
