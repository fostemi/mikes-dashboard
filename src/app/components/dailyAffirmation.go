package components

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func DailyAffirmation() (*fyne.Container, error) {
  rsp, err := http.Get("http://localhost:8080/api/affirmation")
  if err != nil {
    return nil, err
  }
  defer rsp.Body.Close()
  body, err := ioutil.ReadAll(rsp.Body)
  var affirmationRsp AffirmationRsp
  if err := json.Unmarshal(body, &affirmationRsp); err != nil {
    // handle err
    return nil, err
  }

  affirmation := affirmationRsp.Affirmation

  return container.NewVBox(widget.NewLabel("Affirmation of the Day: " + affirmation)), nil
}

type AffirmationRsp struct {
  Affirmation string `json:"affirmation"`
}
