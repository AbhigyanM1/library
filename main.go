package main

import (
    "github.com/gin-gonic/gin"
    "library-api/cassandra"
    "library-api/handlers"
)

func main() {
    cassandra.Init()

    router := gin.Default()
    router.GET("/books", handlers.AddBook)
    router.GET("/users", handlers.AddUser)
    router.GET("/borrow", handlers.BorrowBook)
    router.Run(":8080")
}