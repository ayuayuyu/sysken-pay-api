package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Charge struct {
	ID        uint `gorm:"primaryKey"`
	UserId    uuid.UUID
	Amount    int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
