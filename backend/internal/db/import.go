package db

import (
	"encoding/csv"
	"fmt"
	"github.com/joho/godotenv"
	"githug.com/oliverTuesta/stocks-info/backend/internal/model"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

func ImportCompaniesFromCSV(db *gorm.DB, filePath string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("no .env file found")
	}

	logoBaseURL := os.Getenv("LOGOS_BASE_URL")
	if logoBaseURL == "" {
		return fmt.Errorf("LOGOS_BASE_URL is not set")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// variable columns allowed: no check is performed on the number of fields per record
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read csv: %w", err)
	}

	var companies []model.Company

	for i, row := range records {
		if i == 0 {
			continue
		}

		marketCap, _ := strconv.ParseUint(row[8], 10, 64)
		logoUrl := fmt.Sprintf("%s%s.png", logoBaseURL, row[0])

		company := model.Company{
			Ticker:      row[0],
			CompanyName: row[1],
			ShortName:   row[2],
			Industry:    row[3],
			Description: row[4],
			Website:     row[5],
			Ceo:         row[6],
			Exchange:    row[7],
			MarketCap:   marketCap,
			LogoURL:     logoUrl,
		}

		companies = append(companies, company)
	}

	const batchSize = 100

	for i := 0; i < len(companies); i += batchSize {
		end := min(i+batchSize, len(companies))

		batch := companies[i:end]
		err = db.CreateInBatches(batch, batchSize).Error
		if err != nil {
			return fmt.Errorf("failed to insert stock: %w", err)
		}
	}

	log.Println("Stocks imported successfully")
	return nil
}
