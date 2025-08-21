package usecase

import "github.com/oliverTuesta/stocks-info/backend/internal/domain"

type CompanyUsecase struct {
	store domain.CompanyStore
}

func NewCompanyUsecase(store domain.CompanyStore) *CompanyUsecase {
	return &CompanyUsecase{store: store}
}

func (uc *CompanyUsecase) ListCompanies() ([]domain.Company, error) {
	return uc.store.GetAll()
}
