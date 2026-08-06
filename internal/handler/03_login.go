package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AbdullahBasir/financial-tracker/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type responseBody struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
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

	user, err := cfg.dbQueries.LoginUser(r.Context(), params.Email)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid email or password - Unauthorized: %v", err))
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.PasswordHash)
	if err != nil || !match {
		RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Invalid password - Unauthorized: %v", err))
		return
	}

	expire := 1 * int(time.Hour)
	jwt, err := auth.MakeJWT(user.ID, cfg.JwtSecret, time.Duration(expire))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not make jwt %v", err))
		return
	}

	RespondWithJSON(w, http.StatusOK, responseBody{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		Email:     user.Email,
		Token:     jwt,
	})
}
