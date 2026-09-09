package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"url-shortner/internal/config"
	"url-shortner/internal/handlers"
	"url-shortner/internal/service"
	"url-shortner/internal/storage"
	"url-shortner/internal/middleware"
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

	mux.HandleFunc("GET /", urlHandler.IndexHandler)

	handlerWithLogging := middleware.LoggerMiddleware(mux)

	// Создание и запуск HTTP-сервера
	srv := &http.Server{
		Addr: ":" + strconv.Itoa(cfg.Port),
		Handler: handlerWithLogging,
	}

	//канал сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Force shutdown: %v", err)
	}
	log.Println("Server stopped")

}