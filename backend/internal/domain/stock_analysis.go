package domain

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Price uint64

func (p *Price) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "$", "")
	str = strings.ReplaceAll(str, ",", "")

	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}

	*p = Price(f * 100)
	return nil
}

type StockAnalysis struct {
	ID        uint    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CompanyID uint    `gorm:"column:company_id;not null;index" json:"company_id"`
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CompanyName string    `gorm:"-" json:"company"`
	Ticker      string    `gorm:"column:ticker;not null" json:"ticker"`
	Brokerage   string    `gorm:"column:brokerage;not null" json:"brokerage"`
	Action      string    `gorm:"column:action;not null" json:"action"`
	RatingFrom  string    `gorm:"column:rating_from;not null" json:"rating_from"`
	RatingTo    string    `gorm:"column:rating_to;not null" json:"rating_to"`
	TargetFrom  Price     `gorm:"column:target_from; not null" json:"target_from"`
	TargetTo    Price     `gorm:"column:target_to; not null" json:"target_to"`
	Time        time.Time `gorm:"column:time;not null" json:"time"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"-"`
}
