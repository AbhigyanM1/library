package models

import "github.com/gocql/gocql"

type User struct {
    UserID gocql.UUID `json:"user_id"`
    Name   string     `json:"name"`
}