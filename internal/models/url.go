package models

import (
	"time"
)

//модель данных
type URL struct {
	ID string
	OriginalURL string
	CreatedAt time.Time
}