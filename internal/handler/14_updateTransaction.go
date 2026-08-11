package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) HandlerUpdateTransaction(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Amount      *decimal.Decimal `json:"amount"`
		OccurredAt  *time.Time       `json:"occurred_at"`
		Description *string          `json:"description"`
		AccountID   *string          `json:"account_id"`
		CategoryID  *string          `json:"category_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode request body")
		return
	}

	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}

	transaction, validation := cfg.TransactionValidation(id, r)
	if validation != nil {
		RespondWithError(w, validation.Code, validation.Message)
		return
	}

	updateParams := sqlc.UpdateTransactionParams{
		Amount:      transaction.Amount,
		OccurredAt:  transaction.OccurredAt,
		Description: transaction.Description,
		AccountID:   transaction.AccountID,
		CategoryID:  transaction.CategoryID,
		ID:          transaction.ID,
	}

	if params.Amount != nil {
		if params.Amount.IsZero() {
			RespondWithError(w, http.StatusBadRequest, "amount cannot be zero")
		}
		updateParams.Amount = *params.Amount
	}

	if params.OccurredAt != nil {
		updateParams.OccurredAt = *params.OccurredAt
	}

	if params.Description != nil {
		updateParams.Description = sql.NullString{String: *params.Description, Valid: true}
	}

	if params.AccountID != nil {
		accountID, err := uuid.Parse(*params.AccountID)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid account_id format")
			return
		}
		updateParams.AccountID = accountID
	}

	if params.CategoryID != nil {
		categoryID, err := uuid.Parse(*params.CategoryID)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid category_id format")
			return
		}
		updateParams.CategoryID = categoryID
	}

	updated, err := cfg.dbQueries.UpdateTransaction(r.Context(), updateParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not update transaction")
		return
	}

	RespondWithJSON(w, http.StatusOK, Transaction{
		ID:          updated.ID,
		Amount:      updated.Amount,
		CreatedAt:   updated.CreatedAt,
		OccurredAt:  updated.OccurredAt,
		Description: updated.Description.String,
		AccountID:   updated.AccountID,
		CategoryID:  updated.CategoryID,
	})
}
