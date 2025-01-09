package main

import (
  "os"
  "strings"

  "github.com/joho/godotenv"
)

type User struct {
  ID uint `json:"id"`
  Username string `json:"username"`
  Email string `json:"email"`
  Password string `json:"password"`
}

type SignUpRequest struct {
  Username string `json:"username"`
  Email string `json:"email"`
  Password string `json:"password"`
  ConfirmPassword string `json:"confirmPassword"`
}

type Response struct {
  Results interface{} `json:"results"`
  User User `json:"user"`
  Token string `json:"token" binding:"required"`
}

func getUrl() string {
  _ = godotenv.Load(".env")

  url, ok := os.LookupEnv("API_URL")
  if !ok {
    return "http://localhost:8080"
  }
  return url
}

func trimString(input string) string {
  return strings.TrimSpace(strings.ReplaceAll(input, "\n", ""))
}
