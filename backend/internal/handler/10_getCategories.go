package handler

import (
	"log/slog"
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
		slog.Error("invalid user ID in token subject",
			"error", err,
			"subject", claims.Subject,
		)
		RespondWithError(w, http.StatusUnauthorized, "invalid token subject")
		return
	}

	categories, err := cfg.dbQueries.GetCategories(r.Context(), userID)
	if err != nil {
		slog.Error("failed to retrieve categories from database",
			"error", err,
			"user_id", userID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve categories")
		return
	}

	responseBody := []Category{}
	for _, category := range categories {
		responseBody = append(responseBody, Category{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			Type:      category.Type,
			UserID:    category.UserID,
		})
	}
	RespondWithJSON(w, http.StatusOK, responseBody)
}
