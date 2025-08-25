package usecase

import "github.com/oliverTuesta/stocks-info/backend/internal/domain"

type CompanyUsecase struct {
	repository domain.CompanyRepository
}

func NewCompanyUsecase(repository domain.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{repository: repository}
}

func (uc *CompanyUsecase) ListCompanies(req domain.PaginationRequest) (*domain.PaginatedResult[domain.Company], error) {
	return uc.repository.GetAllPaginated(req)
}

func (uc *CompanyUsecase) GetCompanyByTicker(ticker string) (*domain.Company, error) {
	return uc.repository.GetByTicker(ticker)
}

func (uc *CompanyUsecase) GetHotCompanies(limit int) ([]domain.Company, error) {
	return uc.repository.GetHotCompanies(limit)
}
