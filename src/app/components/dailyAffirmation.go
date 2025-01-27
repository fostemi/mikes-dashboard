package components

import (
	"math/rand"

  "github.com/fostemi/mikes-dashboard/app/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

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
  affirmation := getRandomAffirmation().Affirmation

  return container.NewVBox(widget.NewLabel("Affirmation of the Day: " + affirmation)), nil
}

type Affirmation struct {
  gorm.Model
  ID uint
  Affirmation string `json:"affirmation"`
}

func getAffirmationByID(affirmationID uint) (Affirmation, error) {
  var affirmation Affirmation
  if err := db.DB.Where("id = ?", affirmationID).First(&affirmation).Error; err != nil {
    return affirmation, err
  }
  return affirmation, nil
}

func createAffirmation(msg string) (*Affirmation, error) {
  affirmation := &Affirmation{Affirmation: msg}
  if err := db.DB.Create(&affirmation).Error; err != nil {
    return &Affirmation{}, err
  }
  return affirmation, nil
}
