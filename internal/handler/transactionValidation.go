package handler

import (
	"log/slog"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) TransactionValidation(id string, r *http.Request) (sqlc.Transaction, *ValidationError) {
	transactionID, err := uuid.Parse(id)
	if err != nil {
		return sqlc.Transaction{}, &ValidationError{http.StatusBadRequest, "could not parse transaction ID"}
	}

	checkTransaction, err := cfg.dbQueries.GetTransaction(r.Context(), transactionID)
	if err != nil {
		slog.Error("failed to retrieve transaction from database",
			"error", err,
			"transaction_id", transactionID,
		)
		return sqlc.Transaction{}, &ValidationError{http.StatusNotFound, "transaction not found"}
	}

	claims, ok := r.Context().Value("claims").(*jwt.RegisteredClaims)
	if !ok || claims == nil {
		return sqlc.Transaction{}, &ValidationError{http.StatusUnauthorized, "invalid authentication"}
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		slog.Error("invalid user ID in token subject",
			"error", err,
		)
		return sqlc.Transaction{}, &ValidationError{http.StatusUnauthorized, "invalid token subject"}
	}

	checkAccount, err := cfg.dbQueries.GetAccount(r.Context(), checkTransaction.AccountID)
	if err != nil {
		slog.Error("failed to retrieve account from database",
			"error", err,
			"transaction_id", checkTransaction.AccountID,
		)
		return sqlc.Transaction{}, &ValidationError{http.StatusNotFound, "account not found"}
	}

	if checkAccount.UserID != userID {
		return sqlc.Transaction{}, &ValidationError{http.StatusForbidden, "account does not belong to you"}
	}
	return checkTransaction, nil
}
