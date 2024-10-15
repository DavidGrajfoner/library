package main

import (
	"library/controllers"
	"library/database"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    database.Connect()

    
    r.GET("/users", controllers.GetUsers)
    r.GET("/users/:id", controllers.GetUserByID)
    r.POST("/users", controllers.CreateUser)

    r.GET("/books", controllers.GetBooks)
    r.GET("/books/:id", controllers.GetBookById)
    r.POST("/books", controllers.CreateBook)
    
    r.POST("/borrow", controllers.BorrowBook)
    r.POST("/return", controllers.ReturnBook)


    r.Run(":8080")
}
