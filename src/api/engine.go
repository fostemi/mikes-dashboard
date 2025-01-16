package main

import (
  "net/http"

  "github.com/fostemi/mikes-dashboard/middleware"
  goals "github.com/fostemi/mikes-dashboard/goals"
  "github.com/fostemi/mikes-dashboard/handlers"
  "github.com/fostemi/mikes-dashboard/models"
  "github.com/fostemi/mikes-dashboard/db"

  "github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
  r := gin.Default()

  db.InitDB()
  db.DB.AutoMigrate(&models.User{})

  publicRoutes := r.Group("/api")
  privateRoutes := r.Group("/api")

  publicRoutes.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  publicRoutes.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "healthz",
    })
  })

  publicRoutes.POST("/login", handlers.Login)
  publicRoutes.POST("/signup", handlers.SignUp)

  privateRoutes.Use(middleware.AuthenticationMiddleware())

  privateRoutes.GET("/user", handlers.FindCurrentUser)
  publicRoutes.GET("/user/:id", handlers.FindUser)

  publicRoutes.GET("/affirmation", goals.GetAffirmation)

  // TODO: Implement following endpoints
  // dailyGoals
  privateRoutes.GET("/dailygoal", goals.GetDailyGoals)
  // privateRoutes.POST("/dailygoal")
  // privateRoutes.DELETE("/dailygoal")
  // longtermGoals
  // customWorkout
  // budget
  // schedule

  return r
}
