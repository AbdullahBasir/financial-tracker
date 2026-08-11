package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) HandlerCreateTransaction(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Amount      decimal.Decimal `json:"amount"`
		OccurredAt  time.Time       `json:"occurred_at"`
		Description *string         `json:"description"`
		AccountID   string          `json:"account_id"`
		CategoryID  string          `json:"category_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode request body")
		return
	}

	if params.Amount.IsZero() {
		RespondWithError(w, http.StatusBadRequest, "amount is required")
		return
	}

	if params.OccurredAt.IsZero() {
		RespondWithError(w, http.StatusBadRequest, "occurred_at is required")
		return
	}

	if params.OccurredAt.After(time.Now().UTC()) {
		RespondWithError(w, http.StatusBadRequest, "transaction cannot be in the future")
		return
	}

	claims, ok := r.Context().Value("claims").(*jwt.RegisteredClaims)
	if !ok || claims == nil {
		RespondWithError(w, http.StatusUnauthorized, "invalid authentication")
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		log.Printf("Error parsing user ID from token: %v", err)
		RespondWithError(w, http.StatusUnauthorized, "invalid token subject")
		return
	}

	accountID, err := uuid.Parse(params.AccountID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invalid account_id format")
		return
	}

	categoryID, err := uuid.Parse(params.CategoryID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invalid category_id format")
		return
	}

	account, err := cfg.dbQueries.GetAccount(r.Context(), accountID)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "account not found")
		return
	}

	if account.UserID != userID {
		RespondWithError(w, http.StatusForbidden, "account does not belong to user")
		return
	}

	_, err = cfg.dbQueries.GetCategory(r.Context(), categoryID)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "category not found")
		return
	}

	var description sql.NullString
	if params.Description != nil {
		description = sql.NullString{String: *params.Description, Valid: true}
	}

	transaction, err := cfg.dbQueries.CreateTransaction(r.Context(), sqlc.CreateTransactionParams{
		Amount:      params.Amount,
		OccurredAt:  params.OccurredAt,
		Description: description,
		AccountID:   accountID,
		CategoryID:  categoryID,
	})
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "could not create transaction")
		return
	}

	RespondWithJSON(w, http.StatusCreated, Transaction{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		CreatedAt:   transaction.CreatedAt,
		OccurredAt:  transaction.OccurredAt,
		Description: transaction.Description.String,
		AccountID:   transaction.AccountID,
		CategoryID:  transaction.CategoryID,
	})
}
