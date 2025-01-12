package auth

import (
  "bytes"
  "fmt"
	"log"
	"net/http"
)

func signin() {

  body := []byte(fmt.Sprintf(`{
    "username": %s,
    "password": %s,
  }`, signin.Username.Text, signin.Password.Text))
  resp, err := http.NewRequest("POST", "http://localhost:8080/api/login", bytes.NewBuffer(body))
  if err != nil {
    log.Fatalln(err)
  }
  log.Println(resp)
  log.Println(resp.Body)
  log.Println(resp.Response.Body)

}
