package components

import (
  // "net/http"
  // "encoding/json"
  // "io/ioutil"
  "math/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// func getRandomAffirmation() *models.Affirmation {
// }

func initAffirmations() () {
  
}

func getRandomAffirmation() *Affirmation {
  var affirmations []Affirmation
  affirmations = []Affirmation{
    {Affirmation: "The law of attaction works"},
    {Affirmation: "My beliefs manifest my reality"},
    {Affirmation: "I attract success into my life"},
    {Affirmation: "I have the power to create my reality"},
    {Affirmation: "My thoughts create my reality"},
    {Affirmation: "I believe in the law of attraction"},
    {Affirmation: "I have the power to manifest my dreams"},
    {Affirmation: "I believe deeply that I can achieve anything I desire"},
  }

  return &affirmations[rand.Intn(len(affirmations))]
}

func DailyAffirmation() (*fyne.Container, error) {
  // rsp, err := http.Get("http://localhost:8080/api/affirmation")
  // if err != nil {
  //   return nil, err
  // }
  // defer rsp.Body.Close()
  // body, err := ioutil.ReadAll(rsp.Body)
  // var affirmationRsp AffirmationRsp
  // if err := json.Unmarshal(body, &affirmationRsp); err != nil {
  //   // handle err
  //   return nil, err
  // }

  affirmation := getRandomAffirmation().Affirmation

  return container.NewVBox(widget.NewLabel("Affirmation of the Day: " + affirmation)), nil
}

// type AffirmationRsp struct {
//   Affirmation string `json:"affirmation"`
// }
type Affirmation struct {
  Affirmation string `json:"affirmation"`
}
