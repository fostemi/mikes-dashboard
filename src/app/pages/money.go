package pages

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
)

func MoneyPage() *fyne.Container {
  var _ = widget.NewLabel("")

  moneyContent := container.New(layout.NewGridLayout(2))

  return moneyContent
}
