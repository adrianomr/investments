package models

import "github.com/google/uuid"

type Cdb struct {
	ID             uuid.UUID `json:"id"`
	UserID         string    `json:"user_id"`
	Amount         float64   `json:"amount"`
	CdiPercentage  float64   `json:"cdiPercentage"`
	InvestmentType string    `json:"investmentType"`
}
