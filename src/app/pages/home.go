package pages

import (
  "log"

  "github.com/fostemi/mikes-dashboard/app/components"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

var yearlyGoals = []string{"Build Life Coach App", "Read 24 Books", "Leadville 100", "Two Certifications", "Build 1 Other Project", "Save $3000", "Build Blockchain Gambling Social Media"}
var dailyGoals = []string{"Read Daily Stoic", "Morning Reflection", "Workout", "Read 100 Pages", "Make Dinner", "Evening Reflection", "Commit Code"}

func HomePage() *fyne.Container {
  yearlyGoalsChecklist := components.CheckList(yearlyGoals, widget.NewLabel("2025 Goals"))
  dailyGoalsChecklist := components.CheckList(dailyGoals, widget.NewLabel("Daily Goals"))
  dailyAffirmation, err := components.DailyAffirmation()
  if err != nil {
    log.Fatalln("Error calling affirmation api", err)
  }

  homeContent := container.New(layout.NewGridLayout(2), yearlyGoalsChecklist, dailyGoalsChecklist, dailyAffirmation)

  return homeContent
}
