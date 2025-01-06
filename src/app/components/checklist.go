package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CheckList(list []string, title fyne.Widget) *fyne.Container {
  checkList := container.NewVBox(title)
  for _, item := range list {
    checkList.Add(widget.NewCheck(item, func(bool) {}))
  }

  return checkList
}
