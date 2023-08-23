package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	hasChoice        bool
	projectNameInput textinput.Model
	template         *Template
	templates        map[string]Template
	templateOptions  list.Model
}

func (m *MainModel) sizeTemplateMenuToFrame(msg tea.WindowSizeMsg) {
	_, v := DocStyle.GetFrameSize()
	m.templateOptions.SetSize(msg.Width, msg.Height-v)
}

func (m *MainModel) renderProjectNameUi() string {
	header := "You have chosen to build: " + m.template.Name + " location: " + m.template.Path + "\n\n"
	input := m.projectNameInput.View()
	back := "\n\n<- left arrow to make a different choice"
	return header + input + back
}

func (m *MainModel) makeProjectNameViewReady() (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.hasChoice = true
	cmd = m.projectNameInput.Focus()
	return m, cmd
}

func (m *MainModel) mutateTemplateFromSelected() {
	selectedItem := m.templateOptions.SelectedItem().FilterValue()
	selectedTemplate := m.templates[selectedItem]
	m.template.Name = selectedTemplate.Name
	m.template.Path = selectedTemplate.Path
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key := msg.String(); key == CtrlC || key == Esc {
			return m, tea.Quit
		}
		if key := msg.String(); key == LeftArrow {
			m.hasChoice = false
		}
		if m.hasChoice {
			if key := msg.String(); key == Enter {
				ProjectName = m.projectNameInput.Value()
				return m, tea.Quit
			}
			m.projectNameInput, cmd = m.projectNameInput.Update(msg)
			return m, cmd
		}
		m.templateOptions, cmd = m.templateOptions.Update(msg)
		if key := msg.String(); key == Enter {
			m.mutateTemplateFromSelected()
			return m.makeProjectNameViewReady()
		}
	case tea.WindowSizeMsg:
		m.sizeTemplateMenuToFrame(msg)
	}
	return m, cmd
}

func (m MainModel) View() string {
	if m.hasChoice {
		return m.renderProjectNameUi()
	}
	return m.templateOptions.View()
}

func NewMainModel(template *Template) *MainModel {
	return &MainModel{
		hasChoice:        false,
		projectNameInput: textinput.New(),
		template:         template,
		templateOptions:  CreateTemplateOptionsNavigation(),
		templates:        GetTemplates(),
	}
}
