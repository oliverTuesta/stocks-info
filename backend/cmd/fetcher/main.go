package main

import (
	"fmt"
	"githug.com/oliverTuesta/stocks-info/backend/internal/api"
	"githug.com/oliverTuesta/stocks-info/backend/internal/db"
	"log"
)

func main() {
	fmt.Println("Fetching stock data...")

	stocks, err := api.FetchAllStockRecommendations()
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	conn, err := db.ConnectGorm()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	fmt.Println("Connected to Cockroach DB")

	err = db.Migrate(conn)

	result := conn.Create(&stocks)
	if result.Error != nil {
		log.Fatalf("Failed to insert stock data: %v", result.Error)
	}
}
