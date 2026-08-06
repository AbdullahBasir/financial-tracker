package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerDeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}
	accountID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Error obtaining uuid value from string: %v", err)
		RespondWithError(w, http.StatusBadRequest, "could not parse account ID")
		return
	}

	err = cfg.dbQueries.DeleteAccount(r.Context(), accountID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not delete account")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
