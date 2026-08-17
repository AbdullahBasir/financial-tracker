package handler

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetTransactions(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")
	categoryID := r.URL.Query().Get("category_id")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	pageStr := r.URL.Query().Get("page")

	claims, ok := r.Context().Value("claims").(*jwt.RegisteredClaims)
	if !ok || claims == nil {
		RespondWithError(w, http.StatusUnauthorized, "invalid authentication")
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		slog.Error("invalid user ID in token subject", "error", err)
		RespondWithError(w, http.StatusUnauthorized, "invalid token subject")
		return
	}

	var accountFilter uuid.NullUUID
	if accountID != "" {
		accID, err := uuid.Parse(accountID)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid account_id format")
			return
		}
		accountFilter = uuid.NullUUID{UUID: accID, Valid: true}
	}

	var categoryFilter uuid.NullUUID
	if categoryID != "" {
		catID, err := uuid.Parse(categoryID)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid category_id format")
			return
		}
		categoryFilter = uuid.NullUUID{UUID: catID, Valid: true}
	}

	var fromFilter sql.NullTime
	if from != "" {
		fDate, err := time.Parse("2006-01-02", from)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid from date format (use YYYY-MM-DD)")
			return
		}
		fromFilter = sql.NullTime{Time: fDate, Valid: true}
	}

	var toFilter sql.NullTime
	if to != "" {
		tDate, err := time.Parse("2006-01-02", to)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid to date format (use YYYY-MM-DD)")
			return
		}
		toFilter = sql.NullTime{Time: tDate, Valid: true}
	}

	page := int32(1)
	if pageStr != "" {
		pNum, err := strconv.ParseInt(pageStr, 10, 32)
		if err != nil || pNum < 1 {
			RespondWithError(w, http.StatusBadRequest, "invalid page")
			return
		}
		page = int32(pNum)
	}

	const pageSize = 10
	offset := (page - 1) * int32(pageSize)

	transactions, err := cfg.dbQueries.GetTransactions(r.Context(), sqlc.GetTransactionsParams{
		UserID:     userID,
		AccountID:  accountFilter,
		CategoryID: categoryFilter,
		FromDate:   fromFilter,
		ToDate:     toFilter,
		Limit:      int32(pageSize),
		Offset:     offset,
	})
	if err != nil {
		slog.Error("failed to retrieve transactions from database", "error", err, "user_id", userID)
		RespondWithError(w, http.StatusInternalServerError, "could not fetch transactions")
		return
	}

	result := make([]Transaction, 0, len(transactions))
	for _, t := range transactions {
		var categoryIDPtr *uuid.UUID
		if t.CategoryID.Valid {
			categoryIDPtr = &t.CategoryID.UUID
		}
		result = append(result, Transaction{
			ID:          t.ID,
			Amount:      t.Amount,
			CreatedAt:   t.CreatedAt,
			OccurredAt:  t.OccurredAt,
			Description: t.Description.String,
			AccountID:   t.AccountID,
			CategoryID:  categoryIDPtr,
		})
	}

	RespondWithJSON(w, http.StatusOK, TransactionPage{
		Transaction: result,
		Page:        page,
		PageSize:    pageSize,
	})
}
