package models

import (
  "gorm.io/gorm"
)

type Affirmation struct {
  gorm.Model
  Affirmation string `json:"affirmation"`
}
