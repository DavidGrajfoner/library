package main

import (
	"library/controllers"
	"library/database"
	"log"
	"os"
	"os/signal"
	"syscall"

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


    go func() {
        if err := r.Run(":8080"); err != nil {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()

    gracefulShutdown()
}

func gracefulShutdown() {
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    <-quit

    log.Println("Shutdown signal received, closing database connection...")
    database.Close()
    log.Println("Server shut down gracefully")
}
