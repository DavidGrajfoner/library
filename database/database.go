package database

import (
	"fmt"
	"library/models"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("DB_SSLMODE")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
        host, user, password, dbname, port, sslmode)

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database. DSN: %s, Error: %v", dsn, err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get database object from GORM: %v", err)
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(30 * time.Minute)

    err = DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{})
    if err != nil {
        log.Fatalf("Error during migration: %v", err)
    }

    log.Println("Database connection successfully established!")
}

func Close() {
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get database object for closing: %v", err)
    }

    if err := sqlDB.Close(); err != nil {
        log.Fatalf("Error closing the database connection: %v", err)
    }

    log.Println("Database connection closed successfully.")
}