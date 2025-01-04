package handlers

import "errors"

type LoginRequest struct {
  Username string `json:"username"`
  Email string `json:"email"`
  Password string `json:"password"`
}

func (request *LoginRequest) Validate() error{
  if request.Username == "" && request.Email == "" {
    return errors.New("Username or Email is required")
  }

  return nil
}

type SignupRequest struct {
  Username string `json:"username"`
  Email string `json:"email"`
  Password string `json:"password"`
  ConfirmPassword string `json:"confirmPassword"`
}

func (request *SignupRequest) Validate() error{
  if request.Username == "" && request.Email == "" {
    return errors.New("Username or Email is required")
  }

  if request.Password != request.ConfirmPassword {
    return errors.New("Passwords do not match")
  }

  return nil
}
