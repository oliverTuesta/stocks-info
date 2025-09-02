package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
	"github.com/oliverTuesta/stocks-info/backend/internal/infrastructure/external"
	"gorm.io/gorm"
)

type CompanyUsecase struct {
	repository domain.CompanyRepository
	aiClient   *external.GeminiClient
}

func NewCompanyUsecase(repository domain.CompanyRepository, client *external.GeminiClient) *CompanyUsecase {
	return &CompanyUsecase{
		repository: repository,
		aiClient:   client,
	}
}

func (uc *CompanyUsecase) ListCompanies(req domain.PaginationRequest) (*domain.PaginatedResult[domain.Company], error) {
	return uc.repository.GetAllPaginated(req)
}

func (uc *CompanyUsecase) GetCompanyByTicker(ticker string) (*domain.Company, error) {
	return uc.repository.GetByTicker(ticker)
}

func (uc *CompanyUsecase) GetHotCompanies(limit int) ([]domain.Company, error) {
	return uc.repository.GetHotCompanies(limit)
}

func (uc *CompanyUsecase) GetCompanyAIAnalysis(ticker string) (*domain.AIAnalysis, error) {
	company, err := uc.repository.GetByTicker(ticker)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("company not found")
	}

	analysis, err := uc.repository.GetAIAnalysisByCompanyId(company.ID)
	if err == nil && analysis != nil {
		if time.Since(analysis.UpdatedAt) < 30*24*time.Hour {
			return analysis, nil
		}
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	prompt := generateCompanyAnalysisPrompt(company)
	resp, err := uc.aiClient.GenerateResponse(prompt)
	if err != nil {
		return nil, err
	}

	content := resp.Text()
	log.Printf("AI Response Content: %s", content)
	newAnalysis, err := parseAIAnalysisFromText(content)
	if err != nil {
		return nil, err
	}
	newAnalysis.CompanyID = company.ID
	newAnalysis.AIProvider = uc.aiClient.Model
	newAnalysis.Status = "completed"
	newAnalysis.CreatedAt = time.Now()
	newAnalysis.UpdatedAt = time.Now()

	err = uc.repository.CreateAIAnalysis(newAnalysis)
	if err != nil {
		return nil, err
	}

	return newAnalysis, nil
}

func generateCompanyAnalysisPrompt(company *domain.Company) string {
	contextInfo := fmt.Sprintf(
		"Company Name: %s\nTicker: %s\nIndustry: %s\nSector: %s\nExchange: %s\nMarket Cap: %d\nDescription: %s\nCEO: %s\nWebsite: %s\n",
		company.CompanyName, company.Ticker, company.Industry, company.Sector, company.Exchange, company.MarketCap, company.Description, company.Ceo, company.Website,
	)

	instruction := `You are an expert financial analyst. Produce a concise, factual company analysis and return ONLY a single JSON object (no prose outside the JSON). The JSON must contain the following keys exactly: "investment_considerations", "recent_events", "market_position", "risk_factors", "opportunities", "financial_highlights", "recommendation". Each value should be a short to medium-length string (1-6 sentences). Do not include extra keys.

Example of the required JSON format:
{
  "investment_considerations": "...",
  "recent_events": "...",
  "market_position": "...",
  "risk_factors": "...",
  "opportunities": "...",
  "financial_highlights": "...",
  "recommendation": "..."
}

Return the plain JSON only, without any additional text or formatting. (No markdown, no code blocks, no explanations).`

	return contextInfo + "\n" + instruction
}

func parseAIAnalysisFromText(content string) (*domain.AIAnalysis, error) {
	c := strings.TrimSpace(content)
	start := strings.Index(c, "{")
	end := strings.LastIndex(c, "}")
	c = c[start : end+1]

	var payload struct {
		InvestmentConsiderations string `json:"investment_considerations"`
		RecentEvents             string `json:"recent_events"`
		MarketPosition           string `json:"market_position"`
		RiskFactors              string `json:"risk_factors"`
		Opportunities            string `json:"opportunities"`
		FinancialHighlights      string `json:"financial_highlights"`
		Recommendation           string `json:"recommendation"`
	}

	var aiAnalysis *domain.AIAnalysis = nil
	err := json.Unmarshal([]byte(c), &payload)
	if err != nil {
		return nil, err
	}
	aiAnalysis = &domain.AIAnalysis{
		InvestmentConsiderations: payload.InvestmentConsiderations,
		RecentEvents:             payload.RecentEvents,
		MarketPosition:           payload.MarketPosition,
		RiskFactors:              payload.RiskFactors,
		Opportunities:            payload.Opportunities,
		FinancialHighlights:      payload.FinancialHighlights,
		Recommendation:           payload.Recommendation,
	}
	return aiAnalysis, nil
}
