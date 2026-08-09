package handler

import "net/http"

func (cfg *apiConfig) HandlerDeleteTransaction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) < 1 {
		RespondWithError(w, http.StatusBadRequest, "no id found in request path")
		return
	}

	transaction, validation := cfg.TransactionValidation(id, r)
	if validation != nil {
		RespondWithError(w, validation.Code, validation.Message)
		return
	}

	err := cfg.dbQueries.DeleteTransaction(r.Context(), transaction.ID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "could not delete transaction")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
