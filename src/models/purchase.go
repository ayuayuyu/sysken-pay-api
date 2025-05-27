package models

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uuid.UUID
	ItemId    int
	CreatedAt time.Time
}
