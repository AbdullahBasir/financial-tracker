package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AbdullahBasir/financial-tracker/internal/auth"
	"github.com/AbdullahBasir/financial-tracker/internal/handler"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cfg := handler.LoadConfig()
		tokenString, err := auth.GetBearerToken(r.Header)
		if err != nil {
			handler.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("could not access Bearer token: %v", err))
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
			return []byte(cfg.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			handler.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("unauthorized access: %v", err))
			return
		}

		claims := token.Claims.(*jwt.RegisteredClaims)
		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
