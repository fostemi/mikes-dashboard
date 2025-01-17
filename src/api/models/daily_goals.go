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

func GetCurrentUserDailyGoals(c *gin.Context) (*DailyGoals, error) {
  user, err := GetCurrentUser(c)
  if err != nil {
    return nil, err
  }

  dailyGoals, err := GetDailyGoalsByUser(user)
  if err != nil {
    return nil, err
  }

  return dailyGoals, nil
}

func GetDailyGoalsByUser(user User) (*DailyGoals, error) {
  var dg DailyGoals
  if err := db.DB.Where("user = ?", user).First(&dg).Error; err != nil {
    return nil, err
  }
  return &dg, nil
}

func (dg *DailyGoals) Create() (*DailyGoals, error) {
  if err := db.DB.Create(&dg).Error; err != nil {
    return &DailyGoals{}, err
  }
  return dg, nil
}
