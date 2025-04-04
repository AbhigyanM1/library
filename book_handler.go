package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gocql/gocql"
    "library-api/cassandra"
    "library-api/models"
    "net/http"
)

func AddBook(c *gin.Context) {
    var book models.Book
    if err := c.BindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    book.BookID = gocql.TimeUUID()

    err := cassandra.Session.Query(
        "INSERT INTO books (book_id, title, author, genre) VALUES (?, ?, ?, ?)",
        book.BookID, book.Title, book.Author, book.Genre,
    ).Exec()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
        return
    }

    c.JSON(http.StatusOK, book)
}