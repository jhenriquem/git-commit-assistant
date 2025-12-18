package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modelLoading struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel() modelLoading {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return modelLoading{spinner: s}
}

func (m modelLoading) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m modelLoading) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, nil

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m modelLoading) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Generating... \n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func Loading(stopchan chan bool) {
	p := tea.NewProgram(initialModel())
	for {
		select {
		case <-stopchan:
			p.Kill()
			return
		default:
			if _, err := p.Run(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
