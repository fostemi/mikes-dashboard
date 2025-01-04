package main

import (
  "net/http"

  "github.com/fostemi/mikes-dashboard/middleware"
  "github.com/fostemi/mikes-dashboard/handlers"
  "github.com/fostemi/mikes-dashboard/models"
  "github.com/fostemi/mikes-dashboard/db"

  "github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
  r := gin.Default()

  db.InitDB()
  db.DB.AutoMigrate(&models.User{})

  publicRoute := r.Group("/api")
  privateRoutes := r.Group("/api")


  publicRoute.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  publicRoute.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "healthz",
    })
  })

  publicRoute.POST("/login", handlers.Login)
  publicRoute.POST("/register", handlers.Register)

  privateRoutes.Use(middleware.AuthenticationMiddleware())

  privateRoutes.GET("/user", handlers.FindCurrentUser)
  publicRoute.GET("/user/:id", handlers.FindUser)

  return r
}
