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
    if err := database.DB.WithContext(c.Request.Context()).Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users"})
        return
    }

    var userResponses []user.UserResponse
    for _, u := range users {
        userResponses = append(userResponses, user.UserResponse{
            ID:        u.ID,
            FirstName: u.FirstName,
            LastName:  u.LastName,
        })
    }

    c.JSON(http.StatusOK, userResponses)
}

func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    var u models.User
    if err := database.DB.WithContext(c.Request.Context()).First(&u, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user"})
        return
    }
    
    userResponse := user.UserResponse{
        ID:        u.ID,
        FirstName: u.FirstName,
        LastName:  u.LastName,
    }

    c.JSON(http.StatusOK, userResponse)
}

func CreateUser(c *gin.Context) {
    var request user.CreateUserRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newUser := models.User{
        FirstName: request.FirstName,
        LastName:  request.LastName,
    }

    if err := database.DB.WithContext(c.Request.Context()).Create(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    userResponse := user.UserResponse{
        ID:        newUser.ID,
        FirstName: newUser.FirstName,
        LastName:  newUser.LastName,
    }

    c.JSON(http.StatusCreated, userResponse)
}