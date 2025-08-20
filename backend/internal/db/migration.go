package db

import (
	"githug.com/oliverTuesta/stocks-info/backend/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.StockRecommendation{},
		&model.Company{},
	)
	if err != nil {
		return err
	}
	return nil
}
