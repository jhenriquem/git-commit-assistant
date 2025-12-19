package ui

import (
	lg "github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var StyleError = lg.NewStyle().Foreground(lg.Color("#D8647E")).Bold(true)

var StyleIntroduction = lg.NewStyle().
	Bold(true).
	Foreground(lg.Color("#E15603"))
