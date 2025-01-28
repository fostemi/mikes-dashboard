package pages

import (
	"log"

	"github.com/fostemi/mikes-dashboard/app/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func AffirmationsPage(a fyne.App) *fyne.Container {
  addAffirmationButton := widget.NewButton("Add Affirmation", func() {
    formWindow := a.NewWindow("Add Affirmation")
    newAffirmation := widget.NewEntry()
    form := &widget.Form{
      Items: []*widget.FormItem{
        {Text: "Add Affirmation", Widget: newAffirmation},
      },
      OnSubmit: func() {
        components.CreateAffirmation(newAffirmation.Text)
        formWindow.Close()
      },
    }
    formWindow.SetContent(form)
    formWindow.Show()
  })
  dailyAffirmation, err := components.DailyAffirmation()
  if err != nil {
    log.Fatalln("Error getting affirmations: ", err)
  }

  homeContent := container.New(layout.NewGridLayout(2), dailyAffirmation, addAffirmationButton)

  return homeContent
}
