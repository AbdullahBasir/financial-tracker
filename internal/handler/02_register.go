package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/AbdullahBasir/financial-tracker/internal/auth"
)

func (cfg *apiConfig) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("could not access request body: %v", err))
		return
	}

	if params.Email == "" || params.Password == "" {
		RespondWithError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not hash password: %v", err))
		return
	}

	user, err := cfg.dbQueries.RegisterUser(r.Context(), sqlc.RegisterUserParams{
		Email:        params.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "UNIQUE constraint") {
			RespondWithError(w, http.StatusConflict, "email already registered")
			return
		}
		RespondWithError(w, http.StatusInternalServerError, "could not register user")
		return
	}
	RespondWithJSON(w, http.StatusCreated, User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	})
}
