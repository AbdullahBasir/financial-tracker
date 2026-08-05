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
		Name            string `json:"name"`
		StartingBalance string `json:"starting_balance"`
		Type            string `json:"type"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode request body")
		return
	}

	validTypes := map[string]bool{
		"checking": true,
		"savings":  true,
		"credit":   true,
	}
	if !validTypes[params.Type] {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid account type: %v", err))
		return
	}

	_, err = strconv.ParseFloat(params.StartingBalance, 64)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid starting_balance format: %v", err))
		return
	}

	account, err := cfg.dbQueries.UpdateAccount(r.Context(), sqlc.UpdateAccountParams{
		Name:            params.Name,
		StartingBalance: params.StartingBalance,
		Type:            params.Type,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not update account")
		return
	}

	RespondWithJSON(w, http.StatusCreated, Account{
		ID:              account.ID,
		Name:            account.Name,
		CreatedAt:       account.CreatedAt,
		StartingBalance: account.StartingBalance,
		Type:            account.Type,
		UserID:          account.UserID,
	})
}
