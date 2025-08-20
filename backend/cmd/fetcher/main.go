package main

import (
	"fmt"
	"githug.com/oliverTuesta/stocks-info/backend/internal/db"
	"log"
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

	err = db.ImportCompaniesFromCSV(conn, "data/companies.csv")
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	} else {
		fmt.Println("DB migration completed successfully")
	}

	fmt.Println("Fetching stock data...")

	/*stocks, err := api.FetchAllStockRecommendations()
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	result := conn.Create(&stocks)
	if result.Error != nil {
		log.Fatalf("Failed to insert stock data: %v", result.Error)
	}*/
}
