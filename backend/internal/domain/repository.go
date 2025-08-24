package domain

type CompanyRepository interface {
	GetAll() ([]Company, error)
	FindByTicker(ticker string) (*Company, error)
	CreateMany(company []Company) error
	CreateOne(company *Company) error
}

type StockAnalysisRepository interface {
	CreateMany(analysis []StockAnalysis) error
	CreateOne(analysis *StockAnalysis) error
}
