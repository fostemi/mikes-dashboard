package goals

import (
  "net/http"
  "fmt"

  "github.com/gin-gonic/gin"

  "github.com/fostemi/mikes-dashboard/models"
)

type DailyGoalsRequest struct {
  Goal string `json:"goal"`
}

// GET
// /dailygoals
func GetDailyGoals(ctx *gin.Context) {
  dailyGoals, err := models.GetCurrentUserDailyGoals(ctx)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }
  _ = fmt.Sprintf("")
  goals := dailyGoals.DailyGoals

  ctx.JSON(http.StatusOK, gin.H{"goals": goals})
}

func PostDailyGoals(ctx *gin.Context) {
  var dg DailyGoalsRequest
  if err := ctx.ShouldBindJSON(&dg); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }


  // check if dailygoals exist
  // goals, err := models.GetCurrentUserDailyGoals(ctx)
  // if err != nil {
  //   ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
  // }
  // append to dailyGoals list
  // dailyGoals.Goals.append(dailyGoals.Goals, dg.Goal)
  ctx.JSON(http.StatusOK, gin.H{"Added goal": dg.Goal})
}
