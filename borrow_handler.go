package handlers

import (
    "time"
    "github.com/gin-gonic/gin"
    "library-api/cassandra"
    "library-api/models"
    "net/http"
)

func BorrowBook(c *gin.Context) {
    var record models.BorrowRecord
    if err := c.BindJSON(&record); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    record.BorrowedOn = time.Now()

    err := cassandra.Session.Query(
        "INSERT INTO borrow_records (user_id, book_id, borrowed_on) VALUES (?, ?, ?)",
        record.UserID, record.BookID, record.BorrowedOn,
    ).Exec()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record borrow"})
        return
    }

    c.JSON(http.StatusOK, record)
}