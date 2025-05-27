package repositories

import (
	"sysken-pay/models"

	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, item *models.Item) error {
	return db.Create(item).Error
}
