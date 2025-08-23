package store

import (
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"gorm.io/gorm"
)

type CompanyStoreDB struct {
	db *gorm.DB
}

func NewCompanyStoreDB(db *gorm.DB) *CompanyStoreDB {
	return &CompanyStoreDB{db: db}
}

func (s *CompanyStoreDB) GetAll() ([]domain.Company, error) {
	var companies []domain.Company
	err := s.db.Limit(10).Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil
}
