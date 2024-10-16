package controllers

import (
	"library/database"
	"library/dto/borrow"
	"library/models"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func BorrowBook(c *gin.Context) {
	var req borrow.BorrowReturnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	err := database.DB.WithContext(c.Request.Context()).Transaction(func(tx *gorm.DB) error {
        var user models.User
		if err := tx.First(&user, req.UserID).Error; err != nil {
			utils.HandleError(c, http.StatusNotFound, "User not found", err)
			return err
		}

		var book models.Book
		if err := tx.First(&book, req.BookID).Error; err != nil {
			utils.HandleError(c, http.StatusNotFound, "Book not found", err)
			return err
		}

		if book.Quantity <= 0 {
			utils.HandleError(c, http.StatusBadRequest, "Book not available", nil)
			return gorm.ErrInvalidTransaction
		}

		book.Quantity -= 1
		if err := tx.Save(&book).Error; err != nil {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to update book quantity", err)
			return err
		}

		borrow := models.Borrow{
			UserID: req.UserID,
			BookID: req.BookID,
		}
		if err := tx.Create(&borrow).Error; err != nil {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to create borrow record", err)
			return err
		}

		c.JSON(http.StatusOK, borrow)
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
	}
}


func ReturnBook(c *gin.Context) {
	var req borrow.BorrowReturnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	err := database.DB.WithContext(c.Request.Context()).Transaction(func(tx *gorm.DB) error {
		var book models.Book
		if err := tx.First(&book, req.BookID).Error; err != nil {
			utils.HandleError(c, http.StatusNotFound, "Book not found", err)
			return err
		}

		var borrow models.Borrow
		if err := tx.Where("user_id = ? AND book_id = ?", req.UserID, req.BookID).First(&borrow).Error; err != nil {
			utils.HandleError(c, http.StatusNotFound, "Borrow record not found", err)
			return err
		}

		book.Quantity += 1
		if err := tx.Save(&book).Error; err != nil {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to update book quantity", err)
			return err
		}

		if err := tx.Delete(&borrow).Error; err != nil {
			utils.HandleError(c, http.StatusInternalServerError, "Failed to delete borrow record", err)
			return err
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
	}
}
