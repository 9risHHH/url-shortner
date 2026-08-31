package main

import (
	"log"
	"net/http"
	"strconv"
	"url-shortner/internal/config"
	"url-shortner/internal/handlers"
	"url-shortner/internal/service"
	"url-shortner/internal/storage"
)	

// загрузка конфига
// Настройка маршрутов
func main() {
	cfg := config.Load()

	storage := storage.NewInMemoryStorage() //создаем хранилище
	
	svc := service.NewService(storage) //создаем сервис

	// создаем обработчики
	pingHandler := handlers.PingHandler
	urlHandler := handlers.NewURLHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pingHandler)
	mux.HandleFunc("POST /shorten", urlHandler.ShortenHandler)
	mux.HandleFunc("GET /{id}", urlHandler.RedirectHandler)
	
	// Создание и запуск HTTP-сервера
	srv := &http.Server{
		Addr: ":" + strconv.Itoa(cfg.Port),
		Handler: mux,
	}
	log.Printf("Starting server on port %d", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}