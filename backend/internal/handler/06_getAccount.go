package handler

import (
	"log/slog"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
)

func (cfg *apiConfig) HandlerGetAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}
	account, validation := cfg.AccountValidation(id, r)
	if validation != nil {
		RespondWithError(w, validation.Code, validation.Message)
		return
	}
	expense, err := cfg.dbQueries.GetTotalExpenseAmount(r.Context(), sqlc.GetTotalExpenseAmountParams{
		UserID: account.UserID,
		ID:     account.ID,
	})
	if err != nil {
		slog.Error("failed to retrieve expense from database",
			"error", err,
			"account_id", account.ID,
		)
		RespondWithError(w, http.StatusInternalServerError, "failed to retrieve sum expenses")
	}

	balance := account.StartingBalance.Sub(expense)

	RespondWithJSON(w, http.StatusOK, Account{
		ID:              account.ID,
		Name:            account.Name,
		CreatedAt:       account.CreatedAt,
		StartingBalance: balance,
		Type:            account.Type,
		UserID:          account.UserID,
	})
}
