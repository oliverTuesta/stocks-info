package external

import (
	"encoding/json"
	"fmt"
	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"io"
	"net/http"
)

type StockRecommendationClient struct {
	BaseURL string
	Bearer  string
	Client  *http.Client
}

func NewStockRecommendationClient(baseURL, bearer string) *StockRecommendationClient {
	return &StockRecommendationClient{
		BaseURL: baseURL,
		Bearer:  bearer,
		Client:  http.DefaultClient,
	}
}

func (c *StockRecommendationClient) FetchAll() ([]domain.StockRecommendation, error) {
	var allItems []domain.StockRecommendation
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

		var apiResp domain.StockRecommendationAPIResponse
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
