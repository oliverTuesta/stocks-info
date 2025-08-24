package usecase

import (
	"errors"
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/external"
	"gorm.io/gorm"
)

type AnalysisUsecase struct {
	analysisRepo domain.StockAnalysisRepository
	companyRepo  domain.CompanyRepository
	client       *external.StockAnalysisClient
}

func NewAnalysisUsecase(companyRepo domain.CompanyRepository, analysisRepo domain.StockAnalysisRepository, client *external.StockAnalysisClient) *AnalysisUsecase {
	return &AnalysisUsecase{
		analysisRepo: analysisRepo,
		companyRepo:  companyRepo,
		client:       client,
	}
}

func (uc *AnalysisUsecase) Execute() error {
	items, err := uc.client.FetchAll()
	if err != nil {
		return err
	}
	var analysis = make([]domain.StockAnalysis, 0, len(items))
	for _, item := range items {
		company, err := uc.companyRepo.FindByTicker(item.Ticker)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			continue
		}
		if err != nil {
			return err
		}
		item.CompanyID = company.ID
		analysis = append(analysis, item)
	}
	return uc.analysisRepo.CreateMany(analysis)
}
