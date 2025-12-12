package ui

import (
	"bubbletea-cli/internal/game"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styling
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Padding(0, 1)

	cellStyle = lipgloss.NewStyle().
			Width(5).
			Height(3).
			Align(lipgloss.Center, lipgloss.Center)

	selectedCellStyle = cellStyle.
				Bold(true).
				Background(lipgloss.Color("62"))

	boardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1)

	footerStyle = lipgloss.NewStyle().
			Italic(true).
			Padding(1, 0)
)

// Model
type model struct {
	board    game.Board
	cursor   int // which cell is selected
	turn     int // whose turn game.X or game.O
	msg      string
	quitting bool
	gameOver bool
}

func NewProgram() *tea.Program {
	m := model{board: game.NewBoard(), cursor: 4, turn: game.X}
	p := tea.NewProgram(m, tea.WithAltScreen())
	return p
}

// Init
func (m model) Init() tea.Cmd { return nil }

// ----- Update -----
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			if m.cursor%3 != 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor%3 != 2 {
				m.cursor++
			}
		case "up", "k":
			if m.cursor/3 != 0 {
				m.cursor -= 3
			}
		case "down", "j":
			if m.cursor/3 != 2 {
				m.cursor += 3
			}
		case "enter", " ":
			if m.gameOver {
				return m, nil
			}

			// attempt move
			if err := m.board.MakeMove(m.cursor, m.turn); err != nil {
				m.msg = "Cell occupied — choose another"
				return m, nil
			}
			// check winner
			if winner, ok := m.board.Winner(); ok {
				m.gameOver = true
				if winner == game.X {
					m.msg = "X wins! Press r to restart or q to quit"
				} else {
					m.msg = "O wins! Press r to restart or q to quit"
				}
				return m, nil
			}
			if m.board.Full() {
				m.gameOver = true
				m.msg = "Draw! Press r to restart or q to quit"
				return m, nil
			}
			// toggle turn
			if m.turn == game.X {
				m.turn = game.O
			} else {
				m.turn = game.X
			}
			m.msg = ""
		case "r":
			m.board.Reset()
			m.turn = game.X
			m.gameOver = false
			m.msg = "New game started"
		}
	}
	return m, nil
}

// ----- View helpers -----
func (m model) renderCell(i int) string {
	r := i / 3
	c := i % 3
	val := m.board.Get(r, c)

	text := " "
	switch val {
	case game.X:
		text = "X"
	case game.O:
		text = "O"
	}

	if i == m.cursor {
		return selectedCellStyle.Render(text)
	}
	return cellStyle.Render(text)
}

// ----- View -----
func (m model) View() string {
	if m.quitting {
		return ""
	}

	// Title
	t := titleStyle.Render("Tic-Tac-Toe • Bubble Tea") + "\n\n"

	// Board
	// Board
	rows := []string{}

	for r := 0; r < 3; r++ {
		cells := []string{}
		for c := 0; c < 3; c++ {
			idx := r*3 + c
			cells = append(cells, m.renderCell(idx))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Center, cells...))
	}

	board := boardStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center, rows...),
	)

	// Footer and status
	status := fmt.Sprintf("\n\nTurn: %s | %s\n",
		func() string {
			if m.turn == game.X {
				return "X"
			}
			return "O"
		}(),
		m.msg,
	)
	f := footerStyle.Render("Use arrow keys / h j k l to move • Enter / Space to mark • r to restart • q to quit")

	return t + board + status + "\n" + f

}
