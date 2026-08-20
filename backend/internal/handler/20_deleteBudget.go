package handler

import (
	"log/slog"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/database/sqlc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerDeleteBudget(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}

	BudgetID, err := uuid.Parse(id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "could not parse budget id")
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

	err = cfg.dbQueries.DeleteBudget(r.Context(), sqlc.DeleteBudgetParams{
		ID:     BudgetID,
		UserID: userID,
	})
	if err != nil {
		slog.Error("failed to delete transaction from database",
			"error", err,
			"budget_id", BudgetID,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not delete budget")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
