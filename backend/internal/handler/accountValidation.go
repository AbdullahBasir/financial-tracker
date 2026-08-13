package handler

import (
	"log/slog"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) AccountValidation(id string, r *http.Request) (sqlc.Account, *ValidationError) {
	accountID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Account{}, &ValidationError{http.StatusBadRequest, "could not parse account ID"}
	}

	checkAccount, err := cfg.dbQueries.GetAccount(r.Context(), accountID)
	if err != nil {
		slog.Error("failed to retrieve account from database",
			"error", err,
			"account_id", accountID,
		)
		return sqlc.Account{}, &ValidationError{http.StatusNotFound, "account not found"}
	}

	claims, ok := r.Context().Value("claims").(*jwt.RegisteredClaims)
	if !ok || claims == nil {
		return sqlc.Account{}, &ValidationError{http.StatusUnauthorized, "invalid authentication"}
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		slog.Error("invalid user ID in token subject",
			"error", err,
		)
		return sqlc.Account{}, &ValidationError{http.StatusUnauthorized, "invalid token subject"}
	}

	if checkAccount.UserID != userID {
		return sqlc.Account{}, &ValidationError{http.StatusForbidden, "account does not belong to you"}
	}
	return checkAccount, nil
}
