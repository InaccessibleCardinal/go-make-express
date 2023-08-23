package main

import "github.com/charmbracelet/bubbles/list"

type NavigationItem struct {
	title, description, templateName string
}

func (i NavigationItem) Title() string       { return i.title }
func (i NavigationItem) Description() string { return i.description }
func (i NavigationItem) FilterValue() string { return i.templateName }

func CreateTemplateOptionsNavigation() list.Model {
	navigationItems := []list.Item{
		NavigationItem{
			title:        "Express bare bones project",
			description:  "A bare bones express app. No rules, just right.",
			templateName: "fnBare",
		},
		NavigationItem{
			title:        "Express vanilla functional project",
			description:  "An express app with some optinons on structure: controller funcs, service funcs, some middlware examples.",
			templateName: "fnVan",
		},
		NavigationItem{
			title:        "Express app never throw",
			description:  "Like express vanilla, but with 'neverthrow'.",
			templateName: "fnNt",
		},
		NavigationItem{
			title:        "Express app down with OOP",
			description:  "Like the express vanilla option, but class-based with dependency injection.",
			templateName: "oop",
		},
	}
	nav := list.New(navigationItems, list.NewDefaultDelegate(), 0, 0)
	nav.Title = "Create an express app"
	return nav
}
