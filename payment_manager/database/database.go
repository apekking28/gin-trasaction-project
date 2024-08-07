package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/apekking28/gin_project/payment_manager/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=localhost user=postgres password=postgres dbname=go port=5433 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Transaction{})

    log.Println("Database connected")
}
