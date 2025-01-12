package data

import (
	"encoding/json"
	"io"
	"net/http"
)

type AffirmationRsp struct {
  Affirmation string `json:"affirmation"`
}

func GetAffirmation() (string, error) {
  rsp, err := http.Get("http://localhost:8080/api/affirmation")
  if err != nil {
    return "", err
  }
  defer rsp.Body.Close()
  body, err := io.ReadAll(rsp.Body)
  var affirmationRsp AffirmationRsp
  if err := json.Unmarshal(body, &affirmationRsp); err != nil {
    return "", err
  }

  affirmation := affirmationRsp.Affirmation
  return affirmation, nil
}
