package main

import (
  "strings"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
)

var user string = "mike"

type model struct {}

func (m model) Init() tea.Cmd {
  return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
      return m, tea.Quit
    }
  }
  return m, nil
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).Align(lipgloss.Center)
	contentStyle := lipgloss.NewStyle().Padding(1, 2)

  dashboard := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(strings.Title(user) + "'s Dashboard"),
		contentStyle.Render("Welcome to your personal life dashboard!"),
    "Exit: ctrl+c or q",
  )

  return dashboard
}

