package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID // Identifiant unique
	Username string
	Password string // Hashé (en prod)
	Email    string
}
