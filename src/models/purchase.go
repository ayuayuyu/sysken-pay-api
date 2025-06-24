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

func (purchase *Purchase) Create(userId uuid.UUID, itemId int) {
	purchase.UserId = userId
	purchase.ItemId = itemId
}
