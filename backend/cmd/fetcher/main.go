package main

import (
	"fmt"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/db"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/external"
	"github.com/oliverTuesta/stocks-info/backend/internal/usecase"
	"log"
	"os"
)

func main() {
	conn, err := db.ConnectGorm()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	fmt.Println("Connected to Cockroach DB")

	err = db.Migrate(conn)
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
	err = usecase.ImportCompaniesFromCSV(conn, "data/companies.csv")
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	} else {
		fmt.Println("DB migration completed successfully")
	}

	fmt.Println("Fetching stock data...")

	baseURL := os.Getenv("SOURCE_BASE_URL")
	bearer := os.Getenv("SOURCE_BEARER")

	client := external.NewStockRecommendationClient(baseURL, bearer)
	stocks, err := client.FetchAll()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	result := conn.Create(&stocks)

	if result.Error != nil {
		log.Fatalf("Failed to insert stock data: %v", result.Error)
	}
}
