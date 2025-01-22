package db

import (
  "log/slog"

  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
  slog.Info("Initializing SQLite")
  db, err := gorm.Open(sqlite.Open("var/sqlite.db"), &gorm.Config{})
  if err != nil {
    slog.Error(err.Error())
  }
  DB = db
}
