package pages

import (
	"bytes"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type SignInForm struct {
  Username widget.Entry
  Password widget.Entry
}

func SignInPage(signinWindow, authenticatedWindow fyne.Window) *widget.Form {
  body := []byte(`{
    "username": "Post title",
    "password": "Post description",
    }`)
  var signin SignInForm
  return &widget.Form{
    Items: []*widget.FormItem{
      {Text: "Username", Widget: &signin.Username},
      {Text: "Password", Widget: &signin.Password},
    },
    OnSubmit: func() {
      resp, err := http.NewRequest("POST", "http://localhost:8080/api/login", bytes.NewBuffer(body))
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
      authenticatedWindow.Show()
    },
  }
}
