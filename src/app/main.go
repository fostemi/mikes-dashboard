package main

import (
  "github.com/fostemi/mikes-dashboard/app/pages"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
)


func main() {
  a := app.New()
  w := a.NewWindow("Mikes Dashboard")

  tabs := container.NewAppTabs(
    container.NewTabItem("Home", pages.HomePage()),
    container.NewTabItem("Health", pages.HealthPage()),
    container.NewTabItem("Money", pages.MoneyPage()),
    container.NewTabItem("Schedule", pages.SchedulePage()),
  )

  w.SetMainMenu(fyne.NewMainMenu(
    fyne.NewMenu("Preferences", fyne.NewMenuItem("Preferences", func() {
      w.SetContent(widget.NewLabel("Sign In"))
    })),
  ))
  w.SetContent(tabs)
  w.Resize(fyne.NewSize(700, 500))

  w.ShowAndRun()
}
