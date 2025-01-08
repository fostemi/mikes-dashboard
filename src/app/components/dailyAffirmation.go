package components

import (
  "github.com/fostemi/mikes-dashboard/app/data"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DailyAffirmation() *fyne.Container {
  // TODO: Call an API to get a random affirmation
  affirmation := data.GetAffirmation()
  return container.NewVBox(widget.NewLabel("Affirmation of the Day: " + affirmation))
}
