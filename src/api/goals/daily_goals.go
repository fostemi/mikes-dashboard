package goals

import (
  "net/http"

  "github.com/gin-gonic/gin"

  "github.com/fostemi/mikes-dashboard/models"
)

func GetDailyGoals(ctx *gin.Context) {
  var dailyGoals *models.DailyGoals
  dailyGoals, err := models.GetCurrentUserDailyGoals(ctx)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  ctx.JSON(http.StatusOK, gin.H{"goals": dailyGoals.DailyGoals})
}
