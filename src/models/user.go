package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (user *User) Create(name string) {
	user.ID = uuid.Must(uuid.NewV7())
	user.Name = name
}
