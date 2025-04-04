package models

import (
    "time"
    "github.com/gocql/gocql"
)

type BorrowRecord struct {
    UserID     gocql.UUID `json:"user_id"`
    BookID     gocql.UUID `json:"book_id"`
    BorrowedOn time.Time  `json:"borrowed_on"`
}