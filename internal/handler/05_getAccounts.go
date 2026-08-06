package handler

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerGetAccounts(w http.ResponseWriter, r *http.Request) {
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

	accounts, err := cfg.dbQueries.GetAccounts(r.Context(), userID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not retrieve accounts: %v", err))
		return
	}

	var responseBody []Account
	for _, account := range accounts {
		responseBody = append(responseBody, Account{
			ID:              account.ID,
			Name:            account.Name,
			CreatedAt:       account.CreatedAt,
			StartingBalance: account.StartingBalance,
			Type:            account.Type,
			UserID:          account.UserID,
		})
	}

	sorting := r.URL.Query().Get("sort")
	reverse := sorting == "desc"

	sort.Slice(responseBody, func(i, j int) bool {
		if reverse {
			return responseBody[i].CreatedAt.After(responseBody[j].CreatedAt)
		}
		return responseBody[i].CreatedAt.Before(responseBody[j].CreatedAt)
	})
	RespondWithJSON(w, http.StatusOK, responseBody)
}
