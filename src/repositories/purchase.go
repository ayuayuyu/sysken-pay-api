package repositories

import (
	"sysken-pay/models"

	"gorm.io/gorm"
)

func CreatePurchase(db *gorm.DB, purchase *models.Purchase) error {
	return db.Create(purchase).Error
}
