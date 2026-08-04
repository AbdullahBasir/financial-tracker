package handler

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Account struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	CreatedAt       time.Time `json:"created_at"`
	StartingBalance string    `json:"starting_balance"`
	Type            string    `json:"type"`
	UserID          uuid.UUID `json:"user_id"`
}
