package handler

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetCategories(w http.ResponseWriter, r *http.Request) {
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

	categories, err := cfg.dbQueries.GetCategories(r.Context(), userID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve categories")
		return
	}

	var responseBody []Account
	for _, category := range categories {
		responseBody = append(responseBody, Account{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			Type:      category.Type,
			UserID:    category.UserID,
		})
	}
	RespondWithJSON(w, http.StatusOK, responseBody)
}
