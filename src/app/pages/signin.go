package pages

import (
  "bytes"
  "fmt"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SignInForm struct {
  Username widget.Entry
  Password widget.Entry
}

func SignInPage(client *http.Client, signinWindow, authenticatedWindow fyne.Window) *widget.Form {
  var signin SignInForm
  return &widget.Form{
    Items: []*widget.FormItem{
      {Text: "Username", Widget: &signin.Username},
      {Text: "Password", Widget: &signin.Password},
    },
    OnSubmit: func() {
      body := []byte(fmt.Sprintf(`{
        "username": %s,
        "password": %s,
        }`, signin.Username.Text, signin.Password.Text))
      resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(body))
      // req, err := http.NewRequest("POST", "http://localhost:8080/api/login", bytes.NewBuffer(body))
      if err != nil {
        log.Println("error", err)
      }
      // res, err := client.Do(req)
      //
      log.Println(resp.Body)
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
      authenticatedWindow.Show()
    },
  }
}
