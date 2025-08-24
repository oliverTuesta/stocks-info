package repository

import (
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"gorm.io/gorm"
)

type StockAnalysisRepositoryDB struct {
	db *gorm.DB
}

func NewStockAnalysisRepositoryDB(db *gorm.DB) *StockAnalysisRepositoryDB {
	return &StockAnalysisRepositoryDB{db: db}
}

func (r *StockAnalysisRepositoryDB) CreateMany(analysis []domain.StockAnalysis) error {
	const batchSize = 500
	return r.db.CreateInBatches(&analysis, batchSize).Error
}
func (r *StockAnalysisRepositoryDB) CreateOne(analysis *domain.StockAnalysis) error {
	return r.db.Create(&analysis).Error
}
