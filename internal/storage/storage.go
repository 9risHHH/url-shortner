package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// Storage определяет интерфейс для работы с хранилищем URL
type Storage interface {
	Save(ctx context.Context, originalURL string) (string, error) // сохраняет url и возвращает короткий id
	Get(ctx context.Context, id string) (string, error) // получает оригинальный url по id
}

// InMemoryStorage реализует интерфейс Storage с хранением данных в памяти
type InMemoryStorage struct {
	data map[string]string // key: короткий ID, value: оригинальный URL
	mu sync.RWMutex // защита от конкурентного доступа
}

// Save сохраняет URL в памяти с генерацией уникального ID
func (s *InMemoryStorage) Save(ctx context.Context, originalURL string) (string, error) {
	id := uuid.New().String() // Генерируем уникальный идентификатор

	// Блокируем доступ к map для записи
	s.mu.Lock()
	defer s.mu.Unlock()

	// Сохраняем URL
	s.data[id] = originalURL
	return id, nil
}

// Get возвращает оригинальный URL по ID
// Возвращает ошибку, если URL не найден
func (s *InMemoryStorage) Get(ctx context.Context, id string) (string, error) {
	// Блокируем доступ к map для чтения (другие читатели могут параллельно читать)
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.data[id]
	if !ok {
		return "", fmt.Errorf("url not found %s", id)
	}
	return url, nil
}

// NewInMemoryStorage создает новый экземпляр in-memory хранилища
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string]string),
	}
}