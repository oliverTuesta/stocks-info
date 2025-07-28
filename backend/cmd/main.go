package main

import (
	"fmt"
	"log"

	"githug.com/oliverTuesta/stocks-info/backend/internal/api"
	// "githug.com/oliverTuesta/stocks-info/backend/internal/db"
)

func main() {
	fmt.Println("Fetching stock data...")

	stocks, err := api.FetchAllStockRecommendations()
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	for i := 0; i < len(stocks); i++ {
		fmt.Println(stocks[i].Ticker)
	}

	// conn, err := db.ConnectCockroach()
	// if err != nil {
	// 	log.Fatalf("DB connection failed: %v", err)
	// }
	// defer conn.Close(nil)

	// if err := db.InsertStocks(conn, stocks); err != nil {
	// 	log.Fatalf("Failed to insert: %v", err)
	// }

	fmt.Printf("Successfully stored %d records in DB\n", len(stocks))
}
