package pages

import (
  "image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func FinancesPage() *fyne.Container {
  var _ = widget.NewLabel("")
  // rgba(92, 179, 56, 1)
  green := color.Color()

  budgetHealth := container.New(layout.NewHBoxLayout(), canvas.NewText("Budget Health", canvas.NewText("Healthy", green.RGBA)))


  moneyContent := container.New(layout.NewGridLayout(2))

  return moneyContent
}
