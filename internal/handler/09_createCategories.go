package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not access request body: %v", err))
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

	category, err := cfg.dbQueries.CreateCategory(r.Context(), sqlc.CreateCategoryParams{
		Name:   params.Name,
		Type:   params.Type,
		UserID: userID,
	})
	if err != nil {
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
