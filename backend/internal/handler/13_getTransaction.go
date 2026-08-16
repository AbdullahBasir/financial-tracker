package handler

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetTransaction(w http.ResponseWriter, r *http.Request) {
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

	var categoryIDPtr *uuid.UUID
	if transaction.CategoryID.Valid {
		categoryIDPtr = &transaction.CategoryID.UUID
	}

	RespondWithJSON(w, http.StatusOK, Transaction{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		CreatedAt:   transaction.CreatedAt,
		OccurredAt:  transaction.OccurredAt,
		Description: transaction.Description.String,
		AccountID:   transaction.AccountID,
		CategoryID:  categoryIDPtr,
	})
}
