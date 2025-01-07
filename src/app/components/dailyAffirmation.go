package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DailyAffirmation() *fyne.Container {
  // TODO: Call an API to get a random affirmation
  return container.NewVBox(widget.NewLabel("I am a winner!"))
}
