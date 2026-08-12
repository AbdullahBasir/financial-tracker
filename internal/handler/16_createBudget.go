package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func (cfg *apiConfig) HandlerCreateBudget(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		MonthlyLimit decimal.Decimal `json:"monthly_limit"`
		Month        string          `json:"month"`
		CategoryID   string          `json:"category_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode request body")
		return
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

	categoryID, err := uuid.Parse(params.CategoryID)
	if err != nil {
		log.Printf("Error parsing category ID from request: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "could not parse category id")
		return
	}

	category, err := cfg.dbQueries.GetCategory(r.Context(), categoryID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve category")
		return
	}

	if category.UserID != userID {
		RespondWithError(w, http.StatusForbidden, "category not owned by user")
		return
	}

	budget, err := cfg.dbQueries.CreateBudget(r.Context(), sqlc.CreateBudgetParams{
		MonthlyLimit: params.MonthlyLimit,
		Month:        params.Month,
		UserID:       userID,
		CategoryID:   category.ID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE constraint") {
			RespondWithError(w, http.StatusConflict, "budget already exists")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, "could not create budget")
		return
	}

	RespondWithJSON(w, http.StatusCreated, Budget{
		ID:           budget.ID,
		CreatedAt:    budget.CreatedAt,
		MonthlyLimit: budget.MonthlyLimit,
		Month:        budget.Month,
		UserID:       budget.UserID,
		CategoryID:   budget.CategoryID,
	})
}
