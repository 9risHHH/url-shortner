package handlers

import (
	"net/http"
)

// PingHandler проверяет работоспособность сервера
// GET /ping → возвращает "pong" со статусом 200
// Другие методы → возвращают статус 405 Method Not Allowed
func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong"))
	}
}