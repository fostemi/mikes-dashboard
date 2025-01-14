package utils

import (
  "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
  if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
    return err
  }
  return nil
}
