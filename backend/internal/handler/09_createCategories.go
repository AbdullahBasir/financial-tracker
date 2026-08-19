package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateCategory(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not access request body")
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

	category, err := cfg.dbQueries.CreateCategory(r.Context(), sqlc.CreateCategoryParams{
		Name:   params.Name,
		Type:   params.Type,
		UserID: userID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			RespondWithError(w, http.StatusConflict, "category already exists")
			return
		}
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE constraint") {
			RespondWithError(w, http.StatusConflict, "category already exists")
			return
		}
		slog.Error("failed to create category from database",
			"error", err,
			"user_id", userID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not create category")
		return
	}
	RespondWithJSON(w, http.StatusCreated, Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		Type:      category.Type,
		UserID:    category.UserID,
	})
}
