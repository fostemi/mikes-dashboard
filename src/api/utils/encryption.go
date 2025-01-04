package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
  err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
  if err != nil {
    return err
  }
  return nil
}
