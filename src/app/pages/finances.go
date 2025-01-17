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
  green := color.RGBA{92, 179, 56, 1}

  budgetHealth := container.New(layout.NewHBoxLayout(), canvas.NewText("Budget Health", color.White), layout.NewSpacer(), canvas.NewText("Healthy", green))

  financeContent := container.New(layout.NewVBoxLayout(), budgetHealth)

  return financeContent
}
