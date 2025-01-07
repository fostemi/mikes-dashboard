package main

import (
  "github.com/fostemi/mikes-dashboard/app/components"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

var data = []string{"Goal 1", "Goal 2", "Goal 3"}

func main() {
  a := app.New()
  w := a.NewWindow("Mikes Dashboard")

  checkList := components.CheckList(data, widget.NewLabel("2025 Goals"))
  dailyGoalsList := components.CheckList(data, widget.NewLabel("Daily Goals"))
  dailyAffirmation := components.DailyAffirmation()

  homeContent := container.New(layout.NewGridLayout(2), checkList, dailyGoalsList, dailyAffirmation)
  healthContent := container.New(layout.NewGridLayout(2), checkList, dailyGoalsList, dailyAffirmation)
  moneyContent := container.New(layout.NewGridLayout(2), checkList, dailyGoalsList, dailyAffirmation)
  scheduleContent := container.New(layout.NewGridLayout(2), checkList, dailyGoalsList, dailyAffirmation)

  tabs := container.NewAppTabs(
    container.NewTabItem("Home", homeContent),
    container.NewTabItem("Health", healthContent),
    container.NewTabItem("Money", moneyContent),
    container.NewTabItem("Schedule", scheduleContent),
  )

  preferences := fyne.NewMenuItem("Preferences", func() {
    w.SetContent(widget.NewLabel("Sign In"))
    // To Do: Add a sign in form
    // TODO: add all the preferences and settings for the application
  })

  w.SetMainMenu(fyne.NewMainMenu(
    fyne.NewMenu("Preferences", preferences),
  ))
  w.SetContent(tabs)
  w.Resize(fyne.NewSize(700, 500))

  w.ShowAndRun()
}
