package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerUpdateCategory(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name *string `json:"name"`
		Type *string `json:"type"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not decode request body")
		return
	}

	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}

	CategoryID, err := uuid.Parse(id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not parse category id")
		return
	}

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

	category, err := cfg.dbQueries.GetCategory(r.Context(), CategoryID)
	if err != nil {
		slog.Error("failed to retrieve category from database",
			"error", err,
			"category_id", CategoryID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve category")
		return
	}

	updateParams := sqlc.UpdateCategoryParams{
		Name:   category.Name,
		Type:   category.Type,
		ID:     category.ID,
		UserID: userID,
	}

	if params.Name != nil {
		updateParams.Name = *params.Name
	}
	if params.Type != nil {
		updateParams.Type = *params.Type
	}

	updatedCategory, err := cfg.dbQueries.UpdateCategory(r.Context(), updateParams)
	if err != nil {
		slog.Error("failed to update category from database",
			"error", err,
			"category_id", CategoryID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not update category")
		return
	}
	RespondWithJSON(w, http.StatusOK, Category{
		ID:        updatedCategory.ID,
		Name:      updatedCategory.Name,
		CreatedAt: updatedCategory.CreatedAt,
		Type:      updatedCategory.Type,
		UserID:    updatedCategory.UserID,
	})
}
