package data

import (
  "math/rand"
)

func GetAffirmation() string {
  affirmations := []string{
    "The law of attaction works",
    "My beliefs manifest my reality",
    "I attract success into my life",
    "I have the power to create my reality",
    "My thoughts create my reality",
    "I believe in the law of attraction",
    "I have the power to manifest my dreams",
    "I believe deeply that I can achieve anything I desire",
  }

  return affirmations[rand.Intn(len(affirmations))]
}
