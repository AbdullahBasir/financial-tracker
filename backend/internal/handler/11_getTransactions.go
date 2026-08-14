package handler

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetTransactions(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")
	categoryID := r.URL.Query().Get("category_id")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	pageStr := r.URL.Query().Get("page")

	account, validation := cfg.AccountValidation(accountID, r)
	if validation != nil {
		RespondWithError(w, validation.Code, validation.Message)
		return
	}

	var parsedCategoryID uuid.UUID
	if categoryID != "" {
		catID, err := uuid.Parse(categoryID)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid category_id format")
			return
		}
		parsedCategoryID = catID
	}

	var fromDate time.Time
	if from != "" {
		fDate, err := time.Parse("2006-01-02", from)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid from date format (use YYYY-MM-DD)")
			return
		}
		fromDate = fDate
	}

	var toDate time.Time
	if to != "" {
		tDate, err := time.Parse("2006-01-02", to)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "invalid to date format (use YYYY-MM-DD)")
			return
		}
		toDate = tDate
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
		AccountID: account.ID,
		Limit:     int32(pageSize),
		Offset:    offset,
	})
	if err != nil {
		slog.Error("failed to retrieve transactions from database",
			"error", err,
			"account_id", account.ID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not fetch transactions")
		return
	}

	filtered := []Transaction{}
	for _, transaction := range transactions {
		if categoryID != "" && transaction.CategoryID != parsedCategoryID {
			continue
		}
		if from != "" && transaction.OccurredAt.Before(fromDate) {
			continue
		}
		if to != "" && transaction.OccurredAt.After(toDate) {
			continue
		}
		filtered = append(filtered, Transaction{
			ID:          transaction.ID,
			Amount:      transaction.Amount,
			CreatedAt:   transaction.CreatedAt,
			OccurredAt:  transaction.OccurredAt,
			Description: transaction.Description.String,
			AccountID:   transaction.AccountID,
			CategoryID:  transaction.CategoryID,
		})
	}

	RespondWithJSON(w, http.StatusOK, map[string]any{
		"transactions": filtered,
		"page":         page,
		"page_size":    pageSize,
	})
}
