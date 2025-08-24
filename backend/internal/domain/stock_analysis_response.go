package domain

type StockAnalysisAPIResponse struct {
	Items    []StockAnalysis `json:"items"`
	NextPage string          `json:"next_page"`
}
