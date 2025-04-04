package models

import "github.com/gocql/gocql"

type Book struct {
    BookID gocql.UUID `json:"book_id"`
    Title  string     `json:"title"`
    Author string     `json:"author"`
    Genre  string     `json:"genre"`
}