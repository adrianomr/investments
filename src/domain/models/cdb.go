package models

import "github.com/google/uuid"

type CdbType string

const (
	CdbInvestmentTypeFixed CdbType = "FIXED"
	CdbInvestmentTypeCdi   CdbType = "CDI"
	CdbInvestmentTypeIpca  CdbType = "IPCA"
)

type Cdb struct {
	ID         uuid.UUID `json:"id"`
	UserID     string    `json:"user_id"`
	Amount     float64   `json:"amount"`
	Percentage float64   `json:"percentage" validate:"required"`
	Type       CdbType   `json:"type" validate:"required"`
}
