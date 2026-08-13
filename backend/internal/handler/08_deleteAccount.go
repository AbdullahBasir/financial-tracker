package handler

import (
	"log/slog"
	"net/http"
)

func (cfg *apiConfig) HandlerDeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}
	account, validation := cfg.AccountValidation(id, r)
	if validation != nil {
		RespondWithError(w, validation.Code, validation.Message)
		return
	}

	err := cfg.dbQueries.DeleteAccount(r.Context(), account.ID)
	if err != nil {
		slog.Error("failed to delete account from database",
			"error", err,
		)
		RespondWithError(w, http.StatusInternalServerError, "could not delete account")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
