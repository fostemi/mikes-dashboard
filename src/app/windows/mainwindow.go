package windows

import (
	"github.com/fostemi/mikes-dashboard/app/pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Window interface {
  fyne.Window
}

type MainWindow struct {
  Window fyne.Window
}

func (w MainWindow) CreateMainWindow() {
  tabs := container.NewAppTabs(
    container.NewTabItem("Home", pages.HomePage()),
    // container.NewTabItemWithIcon("Health", icon, pages.HealthPage()),
    container.NewTabItem("Finances", pages.FinancesPage()),
    container.NewTabItem("Health", pages.HealthPage()),
    container.NewTabItem("Education", pages.EducationPage()),
  )
  w.Window.SetMainMenu(fyne.NewMainMenu(
    fyne.NewMenu("Preferences", fyne.NewMenuItem("Preferences", func() {
      w.Window.SetContent(widget.NewLabel("Sign In"))
    })),
  ))
  w.Window.SetContent(tabs)
  w.Window.Resize(fyne.NewSize(700, 500))
  w.Window.SetMaster()
}

func (w MainWindow) Show() {
  w.Window.Show()
}
