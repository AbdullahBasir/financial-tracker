package handler

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetBudgets(w http.ResponseWriter, r *http.Request) {
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

	budgets, err := cfg.dbQueries.GetBudgets(r.Context(), userID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve budgets")
		return
	}

	var response []Budget
	for _, budget := range budgets {
		response = append(response, Budget{
			ID:           budget.ID,
			CreatedAt:    budget.CreatedAt,
			MonthlyLimit: budget.MonthlyLimit,
			Month:        budget.Month,
			UserID:       budget.UserID,
			CategoryID:   budget.CategoryID,
		})
	}
	RespondWithJSON(w, http.StatusOK, response)
}
