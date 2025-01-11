package goals

import (
  "net/http"
  "math/rand"

	"github.com/fostemi/mikes-dashboard/models"
	"github.com/gin-gonic/gin"
)

func getRandomAffirmation() *models.Affirmation {
  var affirmations []models.Affirmation
  affirmations = []models.Affirmation{
    {Affirmation: "The law of attaction works"},
    {Affirmation: "My beliefs manifest my reality"},
    {Affirmation: "I attract success into my life"},
    {Affirmation: "I have the power to create my reality"},
    {Affirmation: "My thoughts create my reality"},
    {Affirmation: "I believe in the law of attraction"},
    {Affirmation: "I have the power to manifest my dreams"},
    models.Affirmation{Affirmation: "I believe deeply that I can achieve anything I desire"},
  }

  return &affirmations[rand.Intn(len(affirmations))]
}

// /api/affirmation
func GetAffirmation(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"affirmation": getRandomAffirmation().Affirmation})
}

