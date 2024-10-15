package controllers

import (
	"library/database"
	"library/dto/borrow"
	"library/models"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


func BorrowBook(c *gin.Context) {
    var req borrow.BorrowReturnRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.HandleError(c, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    var book models.Book
    if err := database.DB.First(&book, req.BookID).Error; err != nil {
        utils.HandleError(c, http.StatusNotFound, "Book not found", err)
        return
    }

    if book.Quantity <= 0 {
        utils.HandleError(c, http.StatusBadRequest, "Book not available", nil)
        return
    }

    book.Quantity -= 1
    database.DB.Save(&book)

    borrow := models.Borrow{UserID: req.UserID, BookID: req.BookID}
    database.DB.Create(&borrow)

    c.JSON(http.StatusOK, borrow)
}

func ReturnBook(c *gin.Context) {
    var req borrow.BorrowReturnRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.HandleError(c, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    var book models.Book
    if err := database.DB.First(&book, req.BookID).Error; err != nil {
        utils.HandleError(c, http.StatusNotFound, "Book not found", err)
        return
    }

    var borrow models.Borrow
    if err := database.DB.Where("user_id = ? AND book_id = ?", req.UserID, req.BookID).First(&borrow).Error; err != nil {
        utils.HandleError(c, http.StatusNotFound, "Borrow record not found", err)
        return
    }

    book.Quantity += 1
    database.DB.Save(&book)

    database.DB.Delete(&borrow)

    c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
