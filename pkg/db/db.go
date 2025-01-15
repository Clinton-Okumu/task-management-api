package db

import (
  "fmt"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "log"
)

var DB *gorm.DB

func ConnectDB() {
  dsn := "host=localhost port=5432 user=clinton password=blindspot dbname=tas_db sslmode=disable"
  var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal("Error connecting to database:", err)
  }
  sqlDB, err := DB.DB
  if err != nil {
    log.Fatalf("Error getting db instance", err)
  }
  if err := sqlDB.Ping(); err != nil {
    log.Fatalf("Error pinging db", err)
  }

  fmt.Println("Connected to database")

}

fun migrateDB() {}
