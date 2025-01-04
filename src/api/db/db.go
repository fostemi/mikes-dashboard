package db

import (
  "os"
  "log/slog"

  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
  vendor := os.Getenv("DB_VENDOR")
  
  switch vendor {
  case "postgres":
    InitPostgres()
  case "mysql":
    InitMySQL()
  case "sqlite":
    InitSQLite()
  default:
    InitSQLite()
  }
}

func InitPostgres() {
  // Initialize Postgres
}

func InitMySQL() {
  // Initialize MySQL
}

func InitSQLite() {
  slog.Info("Initializing SQLite")
  db, err := gorm.Open(sqlite.Open("var/sqlite.db"), &gorm.Config{})
  if err != nil {
    slog.Error("Error initializing SQLite: %v", err)
  }
  DB = db
}
