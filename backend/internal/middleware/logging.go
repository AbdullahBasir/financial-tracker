package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		sWriter := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sWriter, r)

		slog.Info("request handled",
			"method", r.Method,
			"path", r.URL.Path,
			"status", sWriter.status,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

func (sw *statusWriter) WriteHeader(status int) {
	sw.status = status
	sw.ResponseWriter.WriteHeader(status)
}
