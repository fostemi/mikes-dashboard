package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	// "math/rand"
)

func GetAffirmation() (string, error) {
  // affirmations := []string{
  //   "The law of attaction works",
  //   "My beliefs manifest my reality",
  //   "I attract success into my life",
  //   "I have the power to create my reality",
  //   "My thoughts create my reality",
  //   "I believe in the law of attraction",
  //   "I have the power to manifest my dreams",
  //   "I believe deeply that I can achieve anything I desire",
  // }
  //
  // return affirmations[rand.Intn(len(affirmations))]
  rsp, err := http.Get("http://localhost:8080/api/affirmation")
  if err != nil {
    return "", err
  }
  defer rsp.Body.Close()
  body, err := ioutil.ReadAll(rsp.Body)
  var affirmationRsp AffirmationRsp
  if err := json.Unmarshal(body, &affirmationRsp); err != nil {
    // handle err
  }

  affirmation := affirmationRsp.Affirmation
  return affirmation, nil
}

type AffirmationRsp struct {
  Affirmation string `json:"affirmation"`
}
