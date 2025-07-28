package model

type APIResponse struct {
	Items    []StockRecommendation `json:"items"`
	NextPage string                `json:"next_page"`
}
