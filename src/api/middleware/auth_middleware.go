package middleware

import (
  "net/http"
  "fmt"
  "strings"
  
  "github.com/fostemi/mikes-dashboard/utils"
  "github.com/fostemi/mikes-dashboard/models"

  "github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
      c.Abort()
      return
    }
  
    tokenParts := strings.Split(tokenString, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer not parsed correctly"})
      c.Abort()
      return
    }

    tokenString = tokenParts[1]

    claims, err := utils.VerifyToken(tokenString)
    if err != nil {
      fmt.Println("error")
      c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
      c.Abort()
      return
    }

    c.Set("user_id", claims["user_id"])
    c.Next()
  }
}
