package utils

import "golang.org/x/crypto/bcrypt"
import "fmt"

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  fmt.Println(string(bytes))
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
  fmt.Println(password, hash)
  if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
    return err
  }
  return nil
}
