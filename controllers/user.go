package controllers

import (
	"library/database"
	"library/dto/user"
	"library/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
    id := c.Param("id")

    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
    var request user.CreateUserRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := models.User{
        FirstName: request.FirstName,
        LastName:  request.LastName,
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}