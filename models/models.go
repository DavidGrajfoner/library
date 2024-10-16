package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    FirstName string `json:"first_name" gorm:"not null"`
    LastName  string `json:"last_name" gorm:"not null"`
}

type Book struct {
    gorm.Model
    Title    string `json:"title" gorm:"not null"`
    Quantity int    `json:"quantity" gorm:"not null;default:0"`
}

type Borrow struct {
    gorm.Model
    UserID uint `json:"user_id" gorm:"not null;index"`
    BookID uint `json:"book_id" gorm:"not null;index"`
    User   User `json:"user" gorm:"foreignKey:UserID"`
    Book   Book `json:"book" gorm:"foreignKey:BookID"`
}