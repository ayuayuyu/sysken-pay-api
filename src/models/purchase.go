package models

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	ID        int
	UserId    uuid.UUID
	ItemId    int
	CreatedAt time.Time
}
