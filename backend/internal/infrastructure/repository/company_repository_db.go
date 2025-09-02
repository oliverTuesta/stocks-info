package repository

import (
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type CompanyRepositoryDB struct {
	db *gorm.DB
}

func NewCompanyRepositoryDB(db *gorm.DB) *CompanyRepositoryDB {
	return &CompanyRepositoryDB{db: db}
}

func (r *CompanyRepositoryDB) GetAllPaginated(req domain.PaginationRequest) (*domain.PaginatedResult[domain.Company], error) {
	var companies []domain.Company
	var total int64

	// First query for counting (without preloads for better performance)
	countQuery := r.db.Model(&domain.Company{})
	if req.Search != "" {
		searchTerm := "%" + strings.ToLower(req.Search) + "%"
		countQuery = countQuery.Where(
			"LOWER(ticker) LIKE ? OR LOWER(company_name) LIKE ? OR LOWER(short_name) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	// Second query for fetching data with preloads
	query := r.db.Preload("Analyses", func(db *gorm.DB) *gorm.DB {
		return db.Order("time DESC")
	})

	if req.Search != "" {
		searchTerm := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where(
			"LOWER(ticker) LIKE ? OR LOWER(company_name) LIKE ? OR LOWER(short_name) LIKE ?",
			searchTerm, searchTerm, searchTerm,
		)
	}

	offset := (req.Page - 1) * req.PageSize
	if err := query.Limit(req.PageSize).Offset(offset).Find(&companies).Error; err != nil {
		return nil, err
	}

	pagination := domain.PaginationResponse{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
	}
	pagination.CalculateTotalPages()

	return &domain.PaginatedResult[domain.Company]{
		Data:       companies,
		Pagination: pagination,
	}, nil
}

func (r *CompanyRepositoryDB) GetByTicker(ticker string) (*domain.Company, error) {
	var company domain.Company
	ticker = strings.ToUpper(strings.TrimSpace(ticker))

	err := r.db.Preload("Analyses", func(db *gorm.DB) *gorm.DB {
		return db.Order("time DESC, created_at DESC").Limit(1)
	}).Where("ticker = ?", ticker).First(&company).Error

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

func (r *CompanyRepositoryDB) GetHotCompanies(limit int) ([]domain.Company, error) {
	type CompanyWithPercent struct {
		ID         uint    `gorm:"column:id"`
		PercentInc float64 `gorm:"column:percent_inc"`
	}

	var companyResults []CompanyWithPercent

	err := r.db.Table("companies c").
		Select(`c.id, 
            ((sa.target_to - sa.target_from) * 100.0 / sa.target_from) as percent_inc`).
		Joins("JOIN stock_analyses sa ON sa.company_id = c.id").
		Order("percent_inc DESC").
		Limit(limit).
		Scan(&companyResults).Error

	if err != nil {
		return nil, err
	}

	var companyIDs []uint
	for _, result := range companyResults {
		companyIDs = append(companyIDs, result.ID)
	}

	if len(companyIDs) == 0 {
		return []domain.Company{}, nil
	}

	var companies []domain.Company
	err = r.db.Preload("Analyses", func(db *gorm.DB) *gorm.DB {
		return db.Order("time DESC, created_at DESC")
	}).Where("id IN ?", companyIDs).Find(&companies).Error

	if err != nil {
		return nil, err
	}

	companyMap := make(map[uint]domain.Company)
	for _, company := range companies {
		companyMap[company.ID] = company
	}

	var orderedCompanies []domain.Company
	for _, result := range companyResults {
		if company, exists := companyMap[result.ID]; exists {
			orderedCompanies = append(orderedCompanies, company)
		}
	}

	return orderedCompanies, nil
}

func (r *CompanyRepositoryDB) GetAIAnalysisByCompanyId(companyID uint) (*domain.AIAnalysis, error) {
	var analysis domain.AIAnalysis
	err := r.db.Where("company_id = ?", companyID).Order("created_at DESC").Limit(1).First(&analysis).Error
	if err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *CompanyRepositoryDB) CreateAIAnalysis(analysis *domain.AIAnalysis) error {
	return r.db.Create(analysis).Error
}
