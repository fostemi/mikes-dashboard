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
  tableData := [][]string{{"Monthly Budget", "Actual Spend"}, {"$2,300", "$1,000"}}

  budgetHealth := container.New(layout.NewHBoxLayout(), canvas.NewText("Budget Health", color.White), layout.NewSpacer(), canvas.NewText("Healthy", color.White))

  table := widget.NewTable(
    func() (int, int) {
      return len(tableData), len(tableData[0])
    },
    func() fyne.CanvasObject {
      return widget.NewLabel("Monthy Budget")
    },
    func(i widget.TableCellID, o fyne.CanvasObject) {
      o.(*widget.Label).SetText(tableData[i.Row][i.Col])
    },
  )

  financeContent := container.New(layout.NewVBoxLayout(), budgetHealth, table)

  return financeContent
}
