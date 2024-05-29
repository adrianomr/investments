package models

import (
	"github.com/google/uuid"
	"math"
	"time"
)

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

func (c *Cdb) Update(cdis []Cdi, orders []CdbOrder, until time.Time) {
	c.Amount = 0
	if len(orders) == 0 {
		return
	}
	lastOrder := orders[0]
	c.Amount = lastOrder.Amount
	if len(orders) == 1 {
		for _, cdi := range cdis {
			c.Amount = roundFloat(c.Amount+(c.Amount*(cdi.Rate/100)*(c.Percentage/100)), 2)
		}
		return
	}
	for i, order := range orders {
		if i > 0 {
			c.Amount += order.Amount
		}
	}

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
