package handler

import (
	"log/slog"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerDeleteCategory(w http.ResponseWriter, r *http.Request) {
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
		slog.Error("invalid user ID in token subject", "error", err)
		RespondWithError(w, http.StatusUnauthorized, "invalid token subject")
		return
	}

	err = cfg.dbQueries.ArchiveCategory(r.Context(), sqlc.ArchiveCategoryParams{
		ID:     CategoryID,
		UserID: userID,
	})
	if err != nil {
		slog.Error("failed to delete category from database",
			"error", err,
			"category_id", CategoryID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not delete category")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
