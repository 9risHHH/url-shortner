package service

import (
	"context"
	"fmt"
	"strings"
	"url-shortner/internal/storage"
)

type Service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

//валидация URL
func (s *Service) Shorten(ctx context.Context, url string) (string, error) {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "", fmt.Errorf("ошибка: invalid URL format")
	}
	return s.storage.Save(ctx, url)
}

func (s *Service) Resolve(ctx context.Context, id string) (string, error) {
	return s.storage.Get(ctx, id)
}