package usecase

import (
	"strings"
	"testing"

	"github.com/oliverTuesta/stocks-info/backend/internal/domain"
)

func TestParseAIAnalysisFromText_ValidJSON(t *testing.T) {
	jsonContent :=
		`
{
	"investment_considerations": "Consider X.",
	"recent_events": "Event Y.",
	"market_position": "Leader.",
	"risk_factors": "Risks include Z.",
	"opportunities": "Growth in sector.",
	"financial_highlights": "Revenue up 10%.",
	"recommendation": "Hold."
}
`

	res, err := parseAIAnalysisFromText(jsonContent)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if res.InvestmentConsiderations != "Consider X." {
		t.Fatalf("unexpected investment considerations: %s", res.InvestmentConsiderations)
	}
	if res.Recommendation != "Hold." {
		t.Fatalf("unexpected recommendation: %s", res.Recommendation)
	}
}

func TestParseAIAnalysisFromText_WithFencesAndSurroundingText(t *testing.T) {
	content := `
Here is the analysis you requested:
` + "```json" + `
{
	"investment_considerations": "Consider X.",
	"recent_events": "Event Y.",
	"market_position": "Leader.",
	"risk_factors": "Risks include Z.",
	"opportunities": "Growth in sector.",
	"financial_highlights": "Revenue up 10%.",
	"recommendation": "Buy."
}` + "```"

	res, err := parseAIAnalysisFromText(content)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !strings.Contains(res.InvestmentConsiderations, "Consider X") {
		t.Fatalf("investment considerations missing or incorrect: %s", res.InvestmentConsiderations)
	}
	if res.Recommendation != "Buy." {
		t.Fatalf("unexpected recommendation: %s", res.Recommendation)
	}
}

func TestParseAIAnalysisFromText_InvalidInput(t *testing.T) {
	_, err := parseAIAnalysisFromText("no json here")
	if err == nil {
		t.Fatalf("expected error for invalid input, got nil")
	}
}

func TestGenerateCompanyAnalysisPrompt_ContainsKeys(t *testing.T) {
	cmp := &domain.Company{
		CompanyName: "Acme Corp",
		Ticker:      "ACME",
		Industry:    "Tech",
		Sector:      "Software",
		Exchange:    "NASDAQ",
		MarketCap:   123456789,
		Description: "Makes widgets",
		Ceo:         "Jane Doe",
		Website:     "https://acme.example",
	}

	p := generateCompanyAnalysisPrompt(cmp)

	requiredKeys := []string{"investment_considerations", "recent_events", "market_position", "risk_factors", "opportunities", "financial_highlights", "recommendation"}
	for _, k := range requiredKeys {
		if !strings.Contains(p, k) {
			t.Fatalf("prompt missing required key or token: %s", k)
		}
	}

	if !strings.Contains(p, "Return the JSON only") && !strings.Contains(p, "Return ONLY a single JSON object") {
		t.Fatalf("prompt does not instruct to return JSON only")
	}
}
