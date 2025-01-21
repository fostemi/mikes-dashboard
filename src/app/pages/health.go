package pages

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

func HealthPage() *fyne.Container {
  entry := widget.NewEntry()
  form := &widget.Form{
    Items: []*widget.FormItem{
      {Text: "What workout did you do today?", Widget: entry},
    },
    OnSubmit: healthPageSubmit(),
  }
  healthContent := container.New(layout.NewGridLayout(2), form)

  return healthContent
}

func healthPageSubmit() func() {
  return func() {}
}
