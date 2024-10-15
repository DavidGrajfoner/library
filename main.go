package main

import (
	"library/controllers"
	"library/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    err := r.SetTrustedProxies([]string{"192.168.0.0/16", "127.0.0.1"})
    if err != nil {
        log.Fatal("Failed to set trusted proxies:", err)
    }

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
