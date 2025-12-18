package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
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
		tea.Quit()
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
	frames := []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "}
	i := 0
	for {
		select {
		case <-stopchan:

			fmt.Print("\r                          \r")
			fmt.Print("\n")
			return
		default:
			color.RGB(192, 202, 245).Printf("\r  %s Generating... \r ", frames[i%len(frames)])
			time.Sleep(200 * time.Millisecond)
			i++
		}
	}
}
