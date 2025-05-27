package repositories

import (
	"sysken-pay/models"

	"gorm.io/gorm"
)

func CreateCharge(db *gorm.DB, charge *models.Charge) error {
	return db.Create(charge).Error
}
