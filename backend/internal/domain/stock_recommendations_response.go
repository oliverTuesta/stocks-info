package domain

type StockRecommendationAPIResponse struct {
	Items    []StockRecommendation `json:"items"`
	NextPage string                `json:"next_page"`
}
