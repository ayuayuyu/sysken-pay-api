package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Charge struct {
	ID        int
	UserId    uuid.UUID
	Amount    int
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
