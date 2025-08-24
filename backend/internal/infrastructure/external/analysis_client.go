package external

import (
	"encoding/json"
	"fmt"
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"io"
	"net/http"
)

type StockAnalysisClient struct {
	BaseURL string
	Bearer  string
	Client  *http.Client
}

func NewStockAnalysisClient(baseURL, bearer string) *StockAnalysisClient {
	return &StockAnalysisClient{
		BaseURL: baseURL,
		Bearer:  bearer,
		Client:  http.DefaultClient,
	}
}

func (c *StockAnalysisClient) FetchAll() ([]domain.StockAnalysis, error) {
	var allItems []domain.StockAnalysis
	nextPage := ""

	for {
		url := c.BaseURL
		if nextPage != "" {
			url += "?next_page=" + nextPage
		}

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+c.Bearer)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.Client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("request error: %w", err)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var apiResp domain.StockAnalysisAPIResponse
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
