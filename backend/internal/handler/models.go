package handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	CreatedAt       time.Time       `json:"created_at"`
	StartingBalance decimal.Decimal `json:"starting_balance"`
	Type            string          `json:"type"`
	UserID          uuid.UUID       `json:"user_id"`
}

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	UserID    uuid.UUID `json:"user_id"`
}

type Transaction struct {
	ID          uuid.UUID       `json:"id"`
	Amount      decimal.Decimal `json:"amount"`
	CreatedAt   time.Time       `json:"created_at"`
	OccurredAt  time.Time       `json:"occurred_at"`
	Description string          `json:"description"`
	AccountID   uuid.UUID       `json:"account_id"`
	CategoryID  uuid.UUID       `json:"category_id"`
}

type Budget struct {
	ID           uuid.UUID       `json:"id"`
	CreatedAt    time.Time       `json:"created_at"`
	MonthlyLimit decimal.Decimal `json:"monthly_limit"`
	Month        string          `json:"month"`
	UserID       uuid.UUID       `json:"user_id"`
	CategoryID   uuid.UUID       `json:"category_id"`
}

type BudgetSummaryItem struct {
	CategoryID   uuid.UUID       `json:"category_id"`
	CategoryName string          `json:"category_name"`
	MonthlyLimit decimal.Decimal `json:"monthly_limit"`
	TotalSpent   decimal.Decimal `json:"total_spent"`
	Remaining    decimal.Decimal `json:"remaining"`
	IsOverBudget bool            `json:"is_over_budget"`
}

type BudgetSummaryResponse struct {
	Month          string              `json:"month"`
	Items          []BudgetSummaryItem `json:"items"`
	TotalBudget    decimal.Decimal     `json:"total_budget"`
	TotalSpent     decimal.Decimal     `json:"total_spent"`
	TotalRemaining decimal.Decimal     `json:"total_remaining"`
}

type TransactionPage struct {
	Transaction []Transaction `json:"transactions"`
	Page        int32         `json:"page"`
	PageSize    int           `json:"page_size"`
}
