package ui

import (
	"fmt"

	lg "github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var StyleError = lg.NewStyle().Foreground(lg.Color("#D8647E")).Bold(true)

var StyleCommit = lg.NewStyle().
	Bold(true).
	Foreground(lg.Color("#8AA46E"))

var StyleIntroduction = lg.NewStyle().
	Bold(true).
	Foreground(lg.Color("#E15603"))

func Introduction() {
	banner := []string{
		"  ____ ___ _____      _    ____ ____ ___ ____ _____  _    _   _ _____",
		" / ___|_ _|_   _|    / \\  / ___/ ___|_ _/ ___|_   _|/ \\  | \\ | |_   _|",
		"| |  _ | |  | |     / _ \\ \\___ \\___ \\| |\\___ \\ | | / _ \\ |  \\| | | | ",
		"| |_| || |  | |    / ___ \\ ___) |__) | | ___) || |/ ___ \\| |\\  | | |  ",
		" \\____|___| |_|   /_/   \\_\\____/____/___|____/ |_/_/   \\_\\_| \\_| |_|  ",
		"                    I'm an AI-powered Git commit assistant.                 ",
	}

	for _, line := range banner {
		fmt.Println(line)
	}
}
