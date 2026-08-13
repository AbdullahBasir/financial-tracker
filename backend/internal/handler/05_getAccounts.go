package handler

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetAccounts(w http.ResponseWriter, r *http.Request) {
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

	accounts, err := cfg.dbQueries.GetAccounts(r.Context(), userID)
	if err != nil {
		slog.Error("failed to retrieve accounts from database",
			"error", err,
			"user_id", userID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve accounts")
		return
	}

	var responseBody []Account
	for _, account := range accounts {
		responseBody = append(responseBody, Account{
			ID:              account.ID,
			Name:            account.Name,
			CreatedAt:       account.CreatedAt,
			StartingBalance: account.StartingBalance,
			Type:            account.Type,
			UserID:          account.UserID,
		})
	}
	RespondWithJSON(w, http.StatusOK, responseBody)
}
