package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/db"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/external"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/repository"
	"github.com/oliverTuesta/stocks-info/backend/internal/usecase"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	conn, err := db.ConnectGorm()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	fmt.Println("Connected to Cockroach DB")

	err = db.Migrate(conn)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	logoBaseURL := os.Getenv("LOGOS_BASE_URL")
	if logoBaseURL == "" {
		log.Fatal("LOGOS_BASE_URL is not set")
	}

	file, err := os.Open("data/companies.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	compRepo := repository.NewCompanyRepositoryDB(conn)
	compImporter := usecase.NewCompaniesImporter(compRepo, logoBaseURL)
	err = compImporter.Execute(file)

	if err != nil {
		log.Fatal(err)
	}

	baseURL := os.Getenv("SOURCE_BASE_URL")
	bearer := os.Getenv("SOURCE_BEARER")
	client := external.NewStockAnalysisClient(baseURL, bearer)
	stockRepo := repository.NewStockAnalysisRepositoryDB(conn)
	analysisImport := usecase.NewAnalysisUsecase(compRepo, stockRepo, client)
	err = analysisImport.Execute()
}
