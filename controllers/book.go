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
    database.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

func GetBookById(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Coulnd not retrive book"})
		return
	}

	c.JSON(http.StatusOK, book)
}


func CreateBook(c *gin.Context) {
	var request book.CreateBookRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  request.Title,
		Quantity: request.Quantity,

	}

	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}
