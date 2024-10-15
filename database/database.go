package database

import (
	"library/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=db user=user password=password dbname=library port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Migracija modelov
    DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{})
}
