package models

import (
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"

  "github.com/fostemi/mikes-dashboard/db"
)

type DailyGoals struct {
  gorm.Model
  User User `json:"user"`
  DailyGoals []string `json:"dailyGoals"`
}

func GetCurrentUserDailyGoals(c *gin.Context) ([]string, error) {
  user, err := GetCurrentUser(c)
  if err != nil {
    return nil, err
  }
  var dailyGoals DailyGoals
  dailyGoals.GetDailyGoalsByUser(user)

  return dailyGoals.DailyGoals, nil
}

func (dg *DailyGoals) GetDailyGoalsByUser(user User) (error) {
  if err := db.DB.Where("user = ?", user).First(&dg).Error; err != nil {
    return err
  }
  return nil
}

