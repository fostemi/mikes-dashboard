package goals

import (
  "net/http"

	"github.com/fostemi/mikes-dashboard/models"
	"github.com/gin-gonic/gin"
)

func getAffirmation() *models.Affirmation {
  var affirmation models.Affirmation
  return &affirmation
}

// /api/affirmation
func GetRandomAffirmation(ctx *gin.Context)  {
  ctx.JSON(http.StatusOK, gin.H{"affirmation": "Yup"})
}

