/*
Package api fetch all the stock recommendations from the external api and stores into the database
*/
package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"githug.com/oliverTuesta/stocks-info/backend/internal/model"
)

func FetchAllStockRecommendations() ([]model.StockRecommendation, error) {
	var allItems []model.StockRecommendation
	nextPage := ""

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file %w", err)
	}

	url := os.Getenv("SOURCE_BASE_URL")

	for {
		fmt.Println(url)
		if nextPage != "" {
			url += "?next_page=" + nextPage
		}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+os.Getenv("SOURCE_BEARER"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("request error: %w", err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var apiResp model.APIResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			return nil, fmt.Errorf("json parse error: %w", err)
		}

		allItems = append(allItems, apiResp.Items...)

		if apiResp.NextPage == "" {
			break
		}
		nextPage = apiResp.NextPage
	}
	return allItems, nil
}
