package handler

import (
	"encoding/json"
	"log/slog"
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
		slog.Error("invalid user ID in token subject",
			"error", err,
			"subject", claims.Subject,
		)
		RespondWithError(w, http.StatusUnauthorized, "invalid token subject")
		return
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not access request body")
		return
	}

	if params.StartingBalance.LessThan(decimal.New(0, 0)) {
		RespondWithError(w, http.StatusBadRequest, "amount cannot be negative")
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
		slog.Error("failed to create account in database",
			"error", err,
			"user_id", userID,
			"account_name", params.Name,
		)
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
