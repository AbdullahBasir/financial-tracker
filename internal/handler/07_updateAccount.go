package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
)

func (cfg *apiConfig) HandlerUpdateAccount(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name            *string `json:"name"`
		StartingBalance *string `json:"starting_balance"`
		Type            *string `json:"type"`
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
	checkAccount, validation := cfg.AccountValidation(id, r)
	if validation != nil {
		RespondWithError(w, validation.Code, validation.Message)
		return
	}

	validTypes := map[string]bool{
		"checking": true,
		"savings":  true,
		"credit":   true,
	}
	if !validTypes[*params.Type] {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid account type: %v", err))
		return
	}

	updateParams := sqlc.UpdateAccountParams{
		ID:              checkAccount.ID,
		Name:            checkAccount.Name,
		StartingBalance: checkAccount.StartingBalance,
		Type:            checkAccount.Type,
	}

	if params.Name != nil {
		updateParams.Name = *params.Name
	}
	if params.Type != nil {
		updateParams.Type = *params.Type
	}
	if params.StartingBalance != nil {
		_, err := strconv.ParseFloat(*params.StartingBalance, 64)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid starting_balance format")
			return
		}
		updateParams.StartingBalance = *params.StartingBalance
	}

	account, err := cfg.dbQueries.UpdateAccount(r.Context(), updateParams)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not update account")
		return
	}

	RespondWithJSON(w, http.StatusOK, Account{
		ID:              account.ID,
		Name:            account.Name,
		CreatedAt:       account.CreatedAt,
		StartingBalance: account.StartingBalance,
		Type:            account.Type,
		UserID:          account.UserID,
	})
}
