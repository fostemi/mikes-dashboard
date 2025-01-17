package main

import (
	"flag"
	"net/http"

	"github.com/fostemi/mikes-dashboard/app/pages"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
  client := &http.Client{}
  devFlag := flag.Bool("dev", false, "Whether or not to run the signin window")
  flag.Parse()

  a := app.New()
  // Test a new signin feature
  signinWindow := a.NewWindow("Sign In")

  a.Settings().SetTheme(&mikesTheme{})
  w := a.NewWindow("Mikes Dashboard")

  // icon := fyne.NewStaticResource("icons/health.png", nil)
  tabs := container.NewAppTabs(
    container.NewTabItem("Home", pages.HomePage()),
    // container.NewTabItemWithIcon("Health", icon, pages.HealthPage()),
    container.NewTabItem("Finances", pages.FinancesPage()),
    container.NewTabItem("Health", pages.HealthPage()),
    container.NewTabItem("Education", pages.EducationPage()),
  )
  tabs.SetTabLocation(container.TabLocationLeading)

  w.SetMainMenu(fyne.NewMainMenu(
    fyne.NewMenu("Preferences", fyne.NewMenuItem("Preferences", func() {
      w.SetContent(widget.NewLabel("Sign In"))
    })),
  ))
  w.SetContent(tabs)
  w.Resize(fyne.NewSize(700, 500))
  w.SetMaster()

  if (*devFlag) {
    w.Show()
  } else {
    signinWindow.SetContent(pages.SignInPage(client, signinWindow, w))
    signinWindow.Resize(fyne.NewSize(650, 450))
    signinWindow.Show()
  }

  a.Run()
}
