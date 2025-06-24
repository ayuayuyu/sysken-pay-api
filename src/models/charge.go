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

func (charge *Charge) Create(userId uuid.UUID, amount int) {
	charge.UserId = userId
	charge.Amount = amount
}
