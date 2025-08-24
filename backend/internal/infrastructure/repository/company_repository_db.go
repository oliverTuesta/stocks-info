package repository

import (
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CompanyRepositoryDB struct {
	db *gorm.DB
}

func NewCompanyRepositoryDB(db *gorm.DB) *CompanyRepositoryDB {
	return &CompanyRepositoryDB{db: db}
}

func (r *CompanyRepositoryDB) GetAll() ([]domain.Company, error) {
	var companies []domain.Company
	err := r.db.Limit(10).Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *CompanyRepositoryDB) FindByTicker(ticker string) (*domain.Company, error) {
	var company domain.Company
	err := r.db.Where("ticker = ?", ticker).First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepositoryDB) CreateMany(companies []domain.Company) error {
	const batchSize = 500
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(&companies, batchSize).Error
}

func (r *CompanyRepositoryDB) CreateOne(company *domain.Company) error {
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&company).Error
}
