package pages

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

func SchedulePage() *fyne.Container {
  var _ = widget.NewLabel("")

  scheduleContent := container.New(layout.NewGridLayout(2))

  return scheduleContent
}
