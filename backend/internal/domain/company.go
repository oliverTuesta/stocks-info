package domain

import "time"

type Company struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Ticker      string `gorm:"column:ticker;not null;uniqueIndex" json:"ticker"`
	CompanyName string `gorm:"column:company_name;not null" json:"company_name"`
	ShortName   string `gorm:"column:short_name;not null" json:"short_name"`
	Industry    string `gorm:"column:industry;null" json:"industry_name"`
	Description string `gorm:"column:description;null" json:"description"`
	Website     string `gorm:"column:website;null" json:"website"`
	LogoURL     string `gorm:"column:logo_url;null" json:"logo_url"`
	Ceo         string `gorm:"column:ceo;null" json:"ceo"`
	Exchange    string `gorm:"column:exchange;not null" json:"exchange"`
	MarketCap   uint64 `gorm:"column:market_cap;not null" json:"market_cap"`
	Sector      string `gorm:"column:sector;null" json:"sector"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Analyses []StockAnalysis `gorm:"foreignKey:CompanyID" json:"analyses,omitempty"`
}
