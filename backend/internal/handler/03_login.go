package handler

import (
	"encoding/json"
	"log/slog"
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
		RespondWithError(w, http.StatusBadRequest, "could not access request body")
		return
	}

	if params.Email == "" || params.Password == "" {
		RespondWithError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	user, err := cfg.dbQueries.LoginUser(r.Context(), params.Email)
	if err != nil {
		slog.Error("failed to retrieve user from database",
			"error", err,
		)
		RespondWithError(w, http.StatusUnauthorized, "Invalid email or password - Unauthorized")
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.PasswordHash)
	if err != nil {
		slog.Error("failed to check password hash")
		RespondWithError(w, http.StatusInternalServerError, "could not verify password")
		return
	}

	if !match {
		RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	expire := 1 * int(time.Hour)
	jwt, err := auth.MakeJWT(user.ID, cfg.JwtSecret, time.Duration(expire))
	if err != nil {
		slog.Error("failed to generate JWT")
		RespondWithError(w, http.StatusInternalServerError, "could not make jwt")
		return
	}

	RespondWithJSON(w, http.StatusOK, responseBody{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		Email:     user.Email,
		Token:     jwt,
	})
}
