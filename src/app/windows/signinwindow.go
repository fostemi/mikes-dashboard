package windows

import (
  "net/http"

	"github.com/fostemi/mikes-dashboard/app/pages"

	"fyne.io/fyne/v2"
)

type SignInWindow struct {
  Window fyne.Window
  MainWindow fyne.Window
  Client *http.Client
}

func (s SignInWindow) Show() {
    s.Window.SetContent(pages.SignInPage(s.Client, s.Window, s.MainWindow))
    s.Window.Resize(fyne.NewSize(650, 450))
    s.Window.Show()
}
