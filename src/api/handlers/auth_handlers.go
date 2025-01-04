package handlers

import (
  "net/http"
  "strconv"

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
  var signup SignupRequest

  if err := c.ShouldBindJSON(&signup); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
    return
  }

  if err := signup.Validate(); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user := models.User{Username: signup.Username, Email: signup.Email, Password: signup.Password}

  if _, err := user.Create(); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not saave user" + err.Error()})
  }

  token, err := utils.GenerateToken(user.ID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User registerd", "user": user, "token": token})
}

func FindUser(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
    return
  }

  user, err := models.GetUserByID(uint(id))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}
