package graphics

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ice-rider/go_snake/core"
)

type Model struct {
	game *core.Game
	width, height int
}

func NewModel(width, height int) Model {
	game := core.NewGame(width, height)
	game.Start()
	return Model{
		game: game,
		width: width,
		height: height,
	}
}

func NewProgram(width, height int) *tea.Program {
	return tea.NewProgram(NewModel(width, height), tea.WithAltScreen())
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
			return "tick"
		}),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "w", "up":
			m.game.RotateSnake(core.Up)
		case "s", "down":
			m.game.RotateSnake(core.Down)
		case "a", "left":
			m.game.RotateSnake(core.Left)
		case "d", "right":
			m.game.RotateSnake(core.Right)
		}
	case string:
		if msg == "tick" {
			m.game.Next()
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.game.Field.String()
}