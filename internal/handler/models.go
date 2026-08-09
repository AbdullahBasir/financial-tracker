package handler

import (
	"time"

	"github.com/google/uuid"
)

type ValidationError struct {
	Code    int
	Message string
}

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

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	UserID    uuid.UUID `json:"user_id"`
}

type Transaction struct {
	ID          uuid.UUID `json:"id"`
	Amount      string    `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	OccurredAt  time.Time `json:"occured_at"`
	Description string    `json:"description"`
	AccountID   uuid.UUID `json:"account_id"`
	CategoryID  uuid.UUID `json:"category_id"`
}
