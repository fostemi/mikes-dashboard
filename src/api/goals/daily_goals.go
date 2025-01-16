package goals

import (
  "net/http"
  "fmt"

  "github.com/gin-gonic/gin"

  "github.com/fostemi/mikes-dashboard/models"
)

func GetDailyGoals(ctx *gin.Context) {
  dailyGoals, err := models.GetCurrentUserDailyGoals(ctx)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  fmt.Println("Success... so far")

  fmt.Println(dailyGoals)
  ctx.JSON(http.StatusOK, gin.H{"goals": dailyGoals})
}
