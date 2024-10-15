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

    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", controllers.GetUsers)
        userRoutes.GET("/:id", controllers.GetUserByID)
        userRoutes.POST("/", controllers.CreateUser)
    }

    bookRoutes := r.Group("/books")
    {
        bookRoutes.GET("/", controllers.GetBooks)
        bookRoutes.GET("/:id", controllers.GetBookById)
        bookRoutes.POST("/", controllers.CreateBook)
    }
    
    r.POST("/borrow", controllers.BorrowBook)
    r.POST("/return", controllers.ReturnBook)


    r.Run(":8080")
}
