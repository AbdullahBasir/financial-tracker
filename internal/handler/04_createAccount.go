package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name            string    `json:"name"`
		StartingBalance string    `json:"starting_balance"`
		Type            string    `json:"type"`
		UserID          uuid.UUID `json:"user_id"`
	}

	claims := r.Context().Value("claims").(*jwt.RegisteredClaims)
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not convert string to uuid format: %v", err))
		return
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not access request body: %v", err))
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

	account, err := cfg.dbQueries.CreateAccount(r.Context(), sqlc.CreateAccountParams{
		Name:            params.Name,
		StartingBalance: params.StartingBalance,
		Type:            params.Type,
		UserID:          userID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not create account: %v", err))
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
