package domain

type CompanyStore interface {
	GetAll() ([]Company, error)
}
