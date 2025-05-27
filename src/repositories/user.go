package repositories

import (
	"sysken-pay/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}
