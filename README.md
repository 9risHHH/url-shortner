![Go Version](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Completed-brightgreen)

# URL Shortener

A lightweight HTTP service for shortening URLs, written in Go.

## Description

This project is an HTTP service that allows you to shorten long URLs. It uses only the Go standard library and has no external dependencies.

## Technologies

- Go 1.22+
- net/http
- In-memory storage (map + sync.RWMutex)
- Docker

## Installation & Running

### Locally

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/url-shortner.git
cd url-shortner

# Run the application
go mod tidy
go run cmd/app/main.go

Using Makefile

make run       # start the server
make build     # build the binary
make test      # run tests
Using Docker
bash
docker build -t url-shortner .
docker run -p 8080:8080 url-shortner
API
POST /shorten
Create a short URL.

Request:

json
{
  "url": "https://example.com"
}
Response:

json
{
  "short_url": "550e8400-e29b-41d4-a716-446655440000"
}
GET /{id}
Redirect to the original URL.

Example: GET /550e8400-e29b-41d4-a716-446655440000

Response: 302 redirect to https://example.com

GET /ping
Health check endpoint.

Response: pong

Usage Examples

# Create a short URL
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'

# Follow the short URL
curl -v http://localhost:8080/YOUR_SHORT_ID
Project Structure

url-shortner/
├── cmd/app/            # Entry point
├── internal/
│   ├── config/         # Configuration
│   ├── handlers/       # HTTP handlers
│   ├── models/         # Data models
│   ├── service/        # Business logic
│   └── storage/        # In-memory storage
├── migrations/         # SQL migrations (future)
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
Roadmap
□ Short ID generation (6-8 characters)
□ Graceful shutdown
□ PostgreSQL support
□ Request logging
□ Click statistics
□ Web interface

License
MIT


Русская версия
URL Shortener
Сервис для сокращения ссылок на Go.

Описание
HTTP-сервис, который позволяет сокращать длинные URL-адреса. Работает на стандартной библиотеке Go.

Установка и запуск

# Клонирование репозитория
git clone https://github.com/YOUR_USERNAME/url-shortner.git
cd url-shortner

# Запуск приложения
go mod tidy
go run cmd/app/main.go
API
POST /shorten
Создание короткой ссылки.

Запрос:

json
{
  "url": "https://example.com"
}
Ответ:

json
{
  "short_url": "550e8400-e29b-41d4-a716-446655440000"
}
GET /{id}
Редирект на оригинальный URL.

Пример: GET /550e8400-e29b-41d4-a716-446655440000

Ответ: редирект 302 на https://example.com

GET /ping
Проверка работоспособности.

Ответ: pong

Пример использования

# Создание короткой ссылки
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://google.com"}'

# Переход по короткой ссылке
curl -v http://localhost:8080/ВАШ_КОРОТКИЙ_ID
Структура проекта
text
url-shortner/
├── cmd/app/            # Точка входа
├── internal/
│   ├── config/         # Конфигурация
│   ├── handlers/       # HTTP-обработчики
│   ├── models/         # Модели данных
│   ├── service/        # Бизнес-логика
│   └── storage/        # Хранилище
├── migrations/         # SQL-миграции
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
Планы по улучшению
□ Генерация коротких ID (6-8 символов)
□ Graceful shutdown
□ Поддержка PostgreSQL
□ Логирование запросов
□ Статистика переходов
□ Веб-интерфейс

Лицензия
MIT




