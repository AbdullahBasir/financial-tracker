package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name            string          `json:"name"`
		StartingBalance decimal.Decimal `json:"starting_balance"`
		Type            string          `json:"type"`
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

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not access request body: %v", err))
		return
	}

	account, err := cfg.dbQueries.CreateAccount(r.Context(), sqlc.CreateAccountParams{
		Name:            params.Name,
		StartingBalance: params.StartingBalance,
		Type:            params.Type,
		UserID:          userID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE constraint") {
			RespondWithError(w, http.StatusConflict, "account already created")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, "could not create account")
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
