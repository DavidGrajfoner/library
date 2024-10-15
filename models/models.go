package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}

type Book struct {
    gorm.Model
    Title    string `json:"title"`
    Quantity int    `json:"quantity"`
}

type Borrow struct {
    gorm.Model
    UserID uint `json:"user_id"`
    BookID uint `json:"book_id"`
}
