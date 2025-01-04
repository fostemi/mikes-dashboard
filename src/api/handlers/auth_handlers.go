package handlers

import (
  "net/http"

  "github.com/fostemi/mikes-dashboard/models"
  "github.com/fostemi/mikes-dashboard/utils"

  "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
  var login LoginRequest

  if err := c.ShouldBindJSON(&login); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
    return
  }

  if err := login.Validate(); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user, err := models.GetUserByUsername(login.Username, false)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
    return
  }

  match := utils.CheckPasswordHash(login.Password, user.Password)
  if !match {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
    return
  }

  token, err := utils.GenerateToken(user.ID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
    return
  }

  user.Sanitize()
  c.JSON(http.StatusOK, gin.H{"token": token, "msg": "Login successful", "user": user})
}

func Register(c *gin.Context) {
  var user models.User

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
    return
  }

  user.ID = 1
  c.JSON(http.StatusCreated, gin.H{"message": "User registerd", "user_id": user.ID})
}
