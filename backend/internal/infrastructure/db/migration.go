package db

import (
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&domain.StockAnalysis{},
		&domain.Company{},
	)
	if err != nil {
		return err
	}
	return nil
}
