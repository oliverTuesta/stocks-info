package domain

import (
	"time"
)

type AIAnalysis struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	CompanyID uint    `json:"company_id" gorm:"not null;index"`
	Company   Company `json:"company,omitempty" gorm:"foreignKey:CompanyID"`

	InvestmentConsiderations string `json:"investment_considerations" gorm:"type:text"` // What to consider before investing
	RecentEvents             string `json:"recent_events" gorm:"type:text"`             // Relevant recent events
	MarketPosition           string `json:"market_position" gorm:"type:text"`           // Company's market position
	RiskFactors              string `json:"risk_factors" gorm:"type:text"`              // Key risks
	Opportunities            string `json:"opportunities" gorm:"type:text"`             // Growth opportunities
	FinancialHighlights      string `json:"financial_highlights" gorm:"type:text"`      // Financial summary
	Recommendation           string `json:"recommendation" gorm:"type:text"`            // Overall AI recommendation

	AIProvider   string `json:"ai_provider" gorm:"size:50"`
	Status       string `json:"status" gorm:"size:20;default:'pending'"`
	ErrorMessage string `json:"error_message,omitempty" gorm:"type:text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
