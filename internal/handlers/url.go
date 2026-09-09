package handlers

import (
	"url-shortner/internal/service"
	"encoding/json"
	"net/http"
	"url-shortner/internal/handlers/templates"
)

type URLHandler struct {
	service *service.Service
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func NewURLHandler(service *service.Service) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Conten-Type", "text/html; charset=utf-8")
	w.Write([]byte(templates.IndexHTML))
}

func (h *URLHandler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//парсинг JSON
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	//вызов сервиса
	id, err := h.service.Shorten(r.Context(), req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//ответ
	resp := ShortenResponse{ShortURL: id}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

//
func (h *URLHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")

	originalURL, err := h.service.Resolve(r.Context(), id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}



