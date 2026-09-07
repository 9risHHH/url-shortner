package middleware

import (
	"time"
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter {
			ResponseWriter: w,
			status: http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Printf(
			"%s %s %d %s %s",
			r.Method,
			r.URL.Path,
			wrapped.status,
			time.Since(start),
			r.RemoteAddr,
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}