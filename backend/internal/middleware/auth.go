package middleware

import (
	"context"
	"log/slog"
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
			handler.RespondWithError(w, http.StatusUnauthorized, "missing or invalid authorization header")
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
			return []byte(cfg.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			handler.RespondWithError(w, http.StatusUnauthorized, "invalid or expired token")
			return
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			slog.Error("failed to extract claims from valid token")
			handler.RespondWithError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
