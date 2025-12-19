package ui

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type modelInput struct {
	label     string
	textInput textinput.Model
	err       error
}

func Input(input string) string {
	p := tea.NewProgram(initialModelDescription(input))
	model, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	prompt := strings.Split(model.View(), "\n")[1]
	return prompt
}

func initialModelDescription(label string) modelInput {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return modelInput{
		label:     label,
		textInput: ti,
		err:       nil,
	}
}

func (m modelInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m modelInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m modelInput) View() string {
	return fmt.Sprintf("\n\n    %s \n     %s \n",
		m.label,
		m.textInput.View(),
	) + "\n"
}
