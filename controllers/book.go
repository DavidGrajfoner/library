package controllers

import (
	"library/database"
	"library/dto/book"
	"library/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
	if err := database.DB.WithContext(c.Request.Context()).Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve books"})
        return
    }

	var bookResponses []book.BookResponse
    for _, b := range books {
        bookResponses = append(bookResponses, book.BookResponse{
            ID:       b.ID,
            Title:    b.Title,
            Quantity: b.Quantity,
        })
    }

    c.JSON(http.StatusOK, bookResponses)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	var b models.Book
    if err := database.DB.WithContext(c.Request.Context()).First(&b, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve book"})
        return
    }

	bookResponse := book.BookResponse{
        ID:       b.ID,
        Title:    b.Title,
        Quantity: b.Quantity,
    }

	c.JSON(http.StatusOK, bookResponse)
}


func CreateBook(c *gin.Context) {
	var request book.CreateBookRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newBook := models.Book{
        Title:    request.Title,
        Quantity: request.Quantity,
    }

    if err := database.DB.WithContext(c.Request.Context()).Create(&newBook).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create book"})
        return
    }

	bookResponse := book.BookResponse{
        ID:       newBook.ID,
        Title:    newBook.Title,
        Quantity: newBook.Quantity,
    }

	c.JSON(http.StatusCreated, bookResponse)
}
