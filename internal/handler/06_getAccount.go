package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetAccount(w http.ResponseWriter, r *http.Request) {
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

	account, err := cfg.dbQueries.GetAccount(r.Context(), accountID)
	if err != nil {
		log.Printf("Error retrieving account with id: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "could not retrieve account")
		return
	}

	RespondWithJSON(w, http.StatusOK, Account{
		ID:              account.ID,
		Name:            account.Name,
		CreatedAt:       account.CreatedAt,
		StartingBalance: account.StartingBalance,
		Type:            account.Type,
		UserID:          account.UserID,
	})
}
