package main

import "github.com/charmbracelet/lipgloss"

const (
	CtrlC     = "ctrl+c"
	Esc       = "esc"
	LeftArrow = "left"
	Enter     = "enter"
)

var (
	DocStyle    = lipgloss.NewStyle().Margin(1, 2)
	ProjectName = ""
)
