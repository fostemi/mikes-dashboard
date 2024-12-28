package main

import (
  "log"
  tea "github.com/charmbracelet/bubbletea"
)

func main() {
  p := tea.NewProgram(model{})
  if _, err := p.Run(); err != nil {
    log.Fatalf("Error starting Bubble Tea form.")
  }
}
