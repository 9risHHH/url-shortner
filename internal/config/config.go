package config

import (
	"os"
	"strconv"
)
// конфиг хранит настройки приложения
type Config struct {
	Port int // порт хттп-сервера
}

// Load загружает конфигурацию из переменных окружения
// Если PORT не задан, используется значение по умолчанию 8080
func Load() Config {
	var cfg Config

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		portInt = 8080
	}
	cfg.Port = portInt
	return cfg
}