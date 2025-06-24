package repositories

import (
	"sysken-pay/models"

	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB, item *models.Item) error {
	return db.Create(item).Error
}

// ReadItemByJanCode は JanCode を条件にアイテムを1件取得します。
func ReadItemByJanCode(db *gorm.DB, janCode string, item *models.Item) error {
	// JanCode を使ってデータベースからアイテムを検索
	// First はレコードが見つからない場合に gorm.ErrRecordNotFound を返す
	if err := db.Where("jan_code = ?", janCode).First(item).Error; err != nil {
		return err
	}
	return nil
}
