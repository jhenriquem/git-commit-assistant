package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type stopMsg struct{}

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

	case stopMsg:
		return m, tea.Quit

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

func Loading(stopchan chan struct{}) {
	p := tea.NewProgram(initialModel())

	go func() { p.Run() }()
	<-stopchan
	p.Send(stopMsg{})
}
