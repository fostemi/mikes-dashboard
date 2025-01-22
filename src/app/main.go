package main

import (
	"flag"
  "fmt"
	"net/http"

	"github.com/fostemi/mikes-dashboard/app/pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
  devFlag *bool
)

func init() {
  _ = fmt.Sprintf("")
  devFlag = flag.Bool("dev", false, "Whether or not to run the signin window")
  flag.Parse()
}

func main() {
  client := &http.Client{}

  a := app.New()
  a.Settings().SetTheme(&mikesTheme{})

  signinWindow := a.NewWindow("Sign In")
  var w MainWindow
  w.Window = a.NewWindow("Mikes Dashboard")
  w.CreateMainWindow()

  if (*devFlag) {
    w.Window.Show()
  } else {
    signinWindow.SetContent(pages.SignInPage(client, signinWindow, w.Window))
    signinWindow.Resize(fyne.NewSize(650, 450))
    signinWindow.Show()
  }

  a.Run()
}

type Window interface {
  fyne.Window
}
type MainWindow struct {
  Window fyne.Window
}

// func (w MainWindow) InitWindow() {
//   w.Window = a.NewWindow("Mikes Dashboard")
// }

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

func (w MainWindow) SetMainWindowName(name string) {
  w.Window.SetTitle(name)
}
