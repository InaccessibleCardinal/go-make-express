package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func RunApplication(template *Template) {
	model := NewMainModel(template)
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if _, err := p.Run(); err != nil {
		log.Fatalf("error running program %s", err.Error())
	}
}

func main() {
	var selectedTemplate Template
	RunApplication(&selectedTemplate)
	RunCommands(ProjectName, selectedTemplate.Path)
}
