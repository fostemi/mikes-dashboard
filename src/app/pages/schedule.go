package pages

import (
  "github.com/fostemi/mikes-dashboard/app/components"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

func SchedulePage() *fyne.Container {
  var _ = widget.NewLabel("")
  dailyAffirmation := components.DailyAffirmation()

  scheduleContent := container.New(layout.NewGridLayout(2), dailyAffirmation)

  return scheduleContent
}
