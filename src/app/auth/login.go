package auth

import (
)

type LoginResponse struct {
  Msg string `json:"msg"`
  Token string `json:"token"`
  User UserResponse `json:"user"`
}

type UserResponse struct {
  Id uint `json:"id"`
  Username string `json:"username"`
  Email string `json:"email"`
}
