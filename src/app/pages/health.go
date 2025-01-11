package pages

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

func HealthPage() *fyne.Container {
  var _ = widget.NewLabel("")

  healthContent := container.New(layout.NewGridLayout(2))

  return healthContent
}
