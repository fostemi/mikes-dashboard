package main

import (
  "net/http"

  "github.com/fostemi/dashboard/middleware"
  "github.com/gin-gonic/gin"
)

func engine() *gin.Engine {
  r := gin.Default()

  r.Use()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "healthz",
    })
  })

  return r
}
