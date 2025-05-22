package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
