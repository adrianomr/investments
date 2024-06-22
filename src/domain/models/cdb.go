package models

import (
	"github.com/gofiber/fiber/v2/log"
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
	c.Amount = orders[0].Amount
	for i, order := range orders {
		if i > 0 {
			cdisToUpdate, cdisRemaining := findCdisToUpdate(cdis, order)
			cdis = cdisRemaining
			for _, cdi := range cdisToUpdate {
				c.Amount = c.Amount + (c.Amount * (cdi.Rate / 100) * (c.Percentage / 100))
			}
			c.Amount = roundFloat(c.Amount, 2)
			log.Info("amount after cdi: ", c.Amount)
			c.Amount += order.Amount
			log.Info("amount after order: ", c.Amount)
		}
	}
	for _, cdi := range cdis {
		c.Amount = c.Amount + (c.Amount * (cdi.Rate / 100) * (c.Percentage / 100))
	}
	c.Amount = roundFloat(c.Amount, 2)
	log.Info("amount after cdi: ", c.Amount)

}

func findCdisToUpdate(cdis []Cdi, order CdbOrder) ([]Cdi, []Cdi) {
	for i2, cdi := range cdis {
		if isAfterOrEqualsDate(cdi.Date, order.Date) {
			return cdis[:i2], cdis[i2:]
		}
	}
	return nil, nil
}

func isAfterOrEqualsDate(t1, t2 time.Time) bool {
	if t1.Year() > t2.Year() {
		return true
	}
	if t1.Year() < t2.Year() {
		return false
	}
	// Years are equal, compare months
	if t1.Month() > t2.Month() {
		return true
	}
	if t1.Month() < t2.Month() {
		return false
	}
	// Years and months are equal, compare days
	return t1.Day() >= t2.Day()
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
