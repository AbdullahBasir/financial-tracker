package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetSumBudgets(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("claims").(*jwt.RegisteredClaims)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}
	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	month := r.URL.Query().Get("month")
	if month == "" {
		RespondWithError(w, http.StatusBadRequest, "month parameter required")
		return
	}

	if _, err := time.Parse("2006-01", month); err != nil {
		RespondWithError(w, http.StatusBadRequest, "month must be in YYYY-MM format")
		return
	}

	summary, err := cfg.GetBudgetSummary(r.Context(), userID, month)
	if err != nil {
		log.Printf("Error: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "failed to fetch budget summary")
		return
	}
	RespondWithJSON(w, http.StatusOK, summary)
}
