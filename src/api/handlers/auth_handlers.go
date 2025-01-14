package handlers

import (
  "net/http"
  "strconv"
  "fmt"

  "github.com/fostemi/mikes-dashboard/models"
  "github.com/fostemi/mikes-dashboard/utils"
  "github.com/fostemi/mikes-dashboard/db"

  "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
  var login LoginRequest

  if err := c.ShouldBindJSON(&login); err != nil {
    fmt.Println(err)
    fmt.Println("Binding JSON failure")
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := login.Validate(); err != nil {
    fmt.Println("Error in validation")
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user, err := models.GetUserByUsername(login.Username, false)
  if err != nil {
    fmt.Println("cannot find user")
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
    return
  }

  if err := utils.CheckPasswordHash(login.Password, user.Password); err != nil {
    fmt.Println("error checking password")
    c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
    return
  }

  token, err := utils.GenerateToken(user.ID)
  if err != nil {
    fmt.Println("Cannot generate token")
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
    return
  }

  user.Sanitize()
  c.JSON(http.StatusOK, gin.H{"token": token, "msg": "Login successful", "user": user})
}

func SignUp(c *gin.Context) {
  var signup SignupRequest

  if err := c.ShouldBindJSON(&signup); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

  c.JSON(http.StatusCreated, gin.H{"message": "User Signed Up", "user": user, "token": token})
}

func FindCurrentUser(c *gin.Context) {
  user, err := models.GetCurrentUser(c)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
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

func UpdateUser(c *gin.Context) {
  var user models.User

  if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
    return
  }

  var input models.User
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := db.DB.Model(&user).Updates(input).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update user", "error": err.Error()})
  }

  c.JSON(http.StatusOK, gin.H{"message": "User updated", "user": user})
}

func DeleteUser(c *gin.Context) {
  var user models.User

  if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
    return
  }
  db.DB.Delete(&user)
  c.JSON(http.StatusOK, gin.H{"message": "User deleted", "user": user})
}
