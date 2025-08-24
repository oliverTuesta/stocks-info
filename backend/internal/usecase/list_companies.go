package usecase

import "github.com/oliverTuesta/stocks-info/backend/internal/domain"

type ListCompanies struct {
	repository domain.CompanyRepository
}

func NewCompanyUsecase(repository domain.CompanyRepository) *ListCompanies {
	return &ListCompanies{repository: repository}
}

func (uc *ListCompanies) Execute() ([]domain.Company, error) {
	return uc.repository.GetAll()
}
