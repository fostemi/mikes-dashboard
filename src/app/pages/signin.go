package pages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/fostemi/mikes-dashboard/app/auth"
)

type LogInForm struct {
  Username widget.Entry
  Password widget.Entry
}

func SignInPage(client *http.Client, signinWindow, authenticatedWindow fyne.Window) *widget.Form {
  var signin LogInForm
  form := &widget.Form{
    Items: []*widget.FormItem{
      {Text: "Username", Widget: &signin.Username},
      {Text: "Password", Widget: &signin.Password},
    },
    OnSubmit: func() {
      url := "http://localhost:8080/api/login"
      body := []byte(fmt.Sprintf(`{
        "username": "%s",
        "password": "%s"
      }`, signin.Username.Text, signin.Password.Text))

      resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
      if err != nil {
        log.Println("error", err)
      }

      loginResponse := auth.LoginResponse{}
      derr := json.NewDecoder(resp.Body).Decode(&loginResponse)
      if derr != nil {
        log.Println("Error: ", derr)
      }
      if loginResponse.Token != "" {
        signinWindow.Close()
        authenticatedWindow.Show()
      } else {
        log.Println("Unable to login.")
        signinWindow.Content().Refresh()
      }
    },
  }
  return form
}
