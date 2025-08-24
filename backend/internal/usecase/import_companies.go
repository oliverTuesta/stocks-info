package usecase

import (
	"encoding/csv"
	"fmt"
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"io"
	"strconv"
)

type CompaniesImporter struct {
	repo        domain.CompanyRepository
	logoBaseURL string
}

func NewCompaniesImporter(repo domain.CompanyRepository, logoBaseURL string) *CompaniesImporter {
	return &CompaniesImporter{
		repo:        repo,
		logoBaseURL: logoBaseURL,
	}
}

func (uc *CompaniesImporter) Execute(r io.Reader) error {
	reader := csv.NewReader(r)

	// variable columns allowed: no check is performed on the number of fields per record
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read csv: %w", err)
	}

	var companies []domain.Company

	for i, row := range records {
		if i == 0 {
			continue
		}

		marketCap, _ := strconv.ParseUint(row[9], 10, 64)
		logoUrl := fmt.Sprintf("%s%s.png", uc.logoBaseURL, row[6])

		company := domain.Company{
			Ticker:      row[0],
			CompanyName: row[1],
			ShortName:   row[2],
			Industry:    row[3],
			Description: row[4],
			Website:     row[5],
			Ceo:         row[7],
			Exchange:    row[8],
			LogoURL:     logoUrl,
			MarketCap:   marketCap,
			Sector:      row[10],
		}

		companies = append(companies, company)
	}

	return uc.repo.CreateMany(companies)
}
