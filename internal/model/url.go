package model

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OriginalURL string
	ShortCode   string `gorm:"uniqueIndex;size:10"`
	ClickCount  int
	ExpiresAt   *time.Time
	CreatedAt   time.Time
}
