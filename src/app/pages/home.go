package pages

import (
  "log"

  "github.com/fostemi/mikes-dashboard/app/components"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

var data = []string{"Goal 1", "Goal 2", "Goal 3"}

func HomePage() *fyne.Container {
  yearlGoals := components.CheckList(data, widget.NewLabel("2025 Goals"))
  dailyGoals := components.CheckList(data, widget.NewLabel("Daily Goals"))
  dailyAffirmation, err := components.DailyAffirmation()
  if err != nil {
    log.Fatalln("Error calling affirmation api", err)
  }

  homeContent := container.New(layout.NewGridLayout(2), yearlGoals, dailyGoals, dailyAffirmation)

  return homeContent
}
