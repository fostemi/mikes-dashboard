package main

import (
  "github.com/fostemi/mikes-dashboard/app/components"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)

var data = []string{"Goal 1", "Goal 2", "Goal 3"}

func main() {
  a := app.New()
  w := a.NewWindow("Mikes Dashboard")

  title := widget.NewLabel("2025 Goals")
  content := components.CheckList(data, title)

  w.SetContent(content)
  w.Resize(fyne.NewSize(300, 200))

  w.ShowAndRun()
}
