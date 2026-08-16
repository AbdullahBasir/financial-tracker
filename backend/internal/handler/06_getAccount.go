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
	totalExpenses, err := cfg.dbQueries.GetTotalExpenseAmount(r.Context(), sqlc.GetTotalExpenseAmountParams{
		UserID: account.UserID,
		ID:     account.ID,
	})
	if err != nil {
		slog.Error("failed to retrieve expense total", "error", err, "account_id", account.ID)
		RespondWithError(w, http.StatusInternalServerError, "failed to retrieve sum expenses")
		return
	}

	totalIncome, err := cfg.dbQueries.GetTotalIncomeAmount(r.Context(), sqlc.GetTotalIncomeAmountParams{
		UserID: account.UserID,
		ID:     account.ID,
	})
	if err != nil {
		slog.Error("failed to retrieve income total", "error", err, "account_id", account.ID)
		RespondWithError(w, http.StatusInternalServerError, "failed to retrieve sum income")
		return
	}

	balance := account.StartingBalance.Add(totalIncome).Sub(totalExpenses)

	RespondWithJSON(w, http.StatusOK, Account{
		ID:              account.ID,
		Name:            account.Name,
		CreatedAt:       account.CreatedAt,
		StartingBalance: balance,
		Type:            account.Type,
		UserID:          account.UserID,
	})
}
