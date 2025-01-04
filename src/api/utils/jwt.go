package utils

import (
  "time"
  "fmt"
  "strings"
  "net/http"
  "os"
  "strconv"

  "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secretpassword")

func GenerateToken(userID uint) (string, error) {
  claims := jwt.MapClaims{}
  claims["user_id"] = userID
  claims["exp"] = time.Now().Add(time.Hour *1).Unix()

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Invalid signing method")
    }
    return secretKey, nil
  })

  if err != nil {
    return nil, err
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    return claims, nil
  }

  return nil, fmt.Errorf("Invalid token")
}

func ExtractToken(r *http.Request) string {
  keys := r.URL.Query()
  token := keys.Get("Bearer")
  if token != "" {
    return token
  }
  bearerToken := r.Header.Get("Authorization")
  if len(strings.Split(bearerToken, " ")) == 2 {
    return strings.Split(bearerToken, " ")[1]
  }

  return ""
}

func ExtractTokenID(r *http.Request) (uint, error) {
  tokenString := ExtractToken(r)

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("Invalid signing method: %v", token.Header["alg"])
    }
    return []byte(os.Getenv("SECRET_KEY")), nil
  })

  if err != nil {
    return 0, err
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if ok && token.Valid {
    userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
    if err != nil {
      return 0, err
    }
    return uint(userID), nil
  }

  return 0, nil
}
