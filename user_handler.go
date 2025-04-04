package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gocql/gocql"
    "library-api/cassandra"
    "library-api/models"
    "net/http"
)

func AddUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.UserID = gocql.TimeUUID()

    err := cassandra.Session.Query(
        "INSERT INTO users (user_id, name) VALUES (?, ?)",
        user.UserID, user.Name,
    ).Exec()

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
        return
    }

    c.JSON(http.StatusOK, user)
}