package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func FinancesPage() *fyne.Container {
  tableData := [][]string{{"Budget Health", "Monthly Budget", "Actual Spend"}, {"Healthy", "$2,300", "$1,000"}}

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

  financeContent := container.New(layout.NewStackLayout(), table)

  return financeContent
}
