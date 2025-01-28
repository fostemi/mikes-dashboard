package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/fostemi/mikes-dashboard/app/db"
	"github.com/fostemi/mikes-dashboard/app/windows"

	"fyne.io/fyne/v2/app"
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
  db.InitDB()

  a := app.New()
  a.Settings().SetTheme(&mikesTheme{})

  w := windows.MainWindow{
    Window: a.NewWindow("Mikes Dashboard"),
  }
  w.CreateMainWindow(a)

  signin := windows.SignInWindow{
    Window: a.NewWindow("Sign In"),
    MainWindow: w.Window,
    Client: client,
  }

  if (*devFlag) {
    w.Show()
  } else {
    signin.Show()
  }

  a.Run()
}
