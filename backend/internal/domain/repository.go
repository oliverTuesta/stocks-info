package domain

type CompanyRepository interface {
	GetAllPaginated(req PaginationRequest) (*PaginatedResult[Company], error)
	GetByTicker(ticker string) (*Company, error)
	CreateMany(company []Company) error
	CreateOne(company *Company) error
	GetHotCompanies(limit int) ([]Company, error)
	GetAIAnalysisByCompanyId(companyID uint) (*AIAnalysis, error)
	CreateAIAnalysis(analysis *AIAnalysis) error
}

type StockAnalysisRepository interface {
	CreateMany(analysis []StockAnalysis) error
	CreateOne(analysis *StockAnalysis) error
}
