package models

import (
  "net/http"
  "strings"
  "html"

  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
  "github.com/fostemi/mikes-dashboard/utils"
  "github.com/fostemi/mikes-dashboard/db"
)

type User struct {
  gorm.Model
  ID uint `json:"id"`
  Username string `json:"username"`
  Password string `gorm:"type:varchar(60)" json:"password"`
  Email string `json:"email"`
}

func GetCurrentUser(c *gin.Context) (User, error) {
  var user User

  userID, err := utils.ExtractTokenID(c.Request)
  if err != nil {
    return user, err
  }

  user, err = GetUserByID(userID)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
  }

  user.Sanitize()

  return user, nil
}

func GetUserByID(id uint) (User, error) {
  var user User

  if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
    return user, err
  }

  return user, nil
}

func GetUserByUsername(username string, sanitze bool) (User, error) {
  var user User
  if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
    return user, err
  }
  if sanitze {
    user.Sanitize()
  }

  return user, nil
}

func (user *User) Sanitize() {
  user.Password = "++++++++"
}

func (user *User) Create() (*User, error) {
  if err := db.DB.Create(&user).Error; err != nil {
    return &User{}, err
  }

  userDailyGoals := DailyGoals{
    User: *user,
  }

  if err := db.DB.Create(&userDailyGoals).Error; err != nil {
    return user, err
  }

  return user, nil
}

func (user *User) BeforeCreate(*gorm.DB) error {
  password, err := utils.HashPassword(user.Password)
  if err != nil {
    return err
  }

  user.Password = password

  user.Username = html.EscapeString(strings.ToLower(user.Username))
  return nil
}
