package main

import (
  "log"
  "net/http"

  "github.com/fostemi/mikes-dashboard/app/pages"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/container"
)


func main() {
  a := app.New()
  // Test a new signin feature
  signinWindow := a.NewWindow("Sign In")
  signinWidget := widget.NewEntry()

  a.Settings().SetTheme(&mikesTheme{})
  w := a.NewWindow("Mikes Dashboard")

  // icon := fyne.NewStaticResource("icons/health.png", nil)
  tabs := container.NewAppTabs(
    container.NewTabItem("Home", pages.HomePage()),
    // container.NewTabItemWithIcon("Health", icon, pages.HealthPage()),
    container.NewTabItem("Health", pages.HealthPage()),
    container.NewTabItem("Money", pages.MoneyPage()),
    container.NewTabItem("Schedule", pages.SchedulePage()),
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

  form := &widget.Form{
    Items: []*widget.FormItem{
      {Text: "SignIn", Widget: signinWidget},
    },
    OnSubmit: func() {
      // test request
      resp, err := http.Get("http://localhost:8080/api/affirmation")
      if err != nil {
        log.Fatalln(err)
      }
      log.Println(resp)

      // Authenticate
      // make sure user is authenticated
      // store JWT token
      // if token 
      // signinWindow.Close()
      // w.Show()
      // else, 
      // handle failed login
      
      signinWindow.Close()
      // setupWindow()
      w.Show()
    },
  }
  signinWindow.SetContent(form)
  signinWindow.Resize(fyne.NewSize(650, 450))
  signinWindow.Show()

  a.Run()
}
