package main

import (
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

type model struct {}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	contentStyle := lipgloss.NewStyle().Padding(1, 2)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("Dashboard"),
		contentStyle.Render("Welcome to your Go-powered terminal dashboard!"),
	)
}

