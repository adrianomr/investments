package models

import (
	"github.com/adrianomr/investments/src/domain/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCdbUpdate(t *testing.T) {
	cdb := &models.Cdb{
		UserID:     "123",
		Amount:     0,
		Percentage: 100,
		Type:       models.CdbInvestmentTypeCdi,
	}

	t.Run("Should update CDB amount when no orders created", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{}, time.Now())

		assert.Equal(t, float64(0), cdb.Amount)
	})

	t.Run("Should update CDB amount when one order created", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{models.CdbOrder{
			UserID: "123",
			Type:   models.CdbOrderTypeBuy,
			Amount: 155.55,
		}}, time.Now())

		assert.Equal(t, 155.55, cdb.Amount)
	})

	t.Run("Should update CDB amount when first buy is today", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{models.CdbOrder{
			UserID: "123",
			Type:   models.CdbOrderTypeBuy,
			Amount: 155.55,
			Date:   time.Now(),
		}}, time.Now())

		assert.Equal(t, 155.55, cdb.Amount)
	})

	t.Run("Should update CDB amount when two purchases created and no CDI info", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{
			{
				UserID: "123",
				Type:   models.CdbOrderTypeBuy,
				Amount: 155.55,
				Date:   time.Now(),
			},
			{
				UserID: "123",
				Type:   models.CdbOrderTypeBuy,
				Amount: 155.55,
				Date:   time.Now(),
			}}, time.Now())

		assert.Equal(t, 311.10, cdb.Amount)
	})

	t.Run("Should update CDB amount when position is next month", func(t *testing.T) {
		cdb.Update([]models.Cdi{models.Cdi{
			Rate: 0.68591,
			Date: time.Date(2024, 05, 24, 12, 12, 12, 0, time.UTC),
		}}, []models.CdbOrder{models.CdbOrder{
			UserID: "123",
			Type:   models.CdbOrderTypeBuy,
			Amount: 90940.38,
			Date:   time.Now(),
		}}, time.Date(2024, 05, 01, 12, 12, 12, 0, time.UTC))

		assert.Equal(t, 91564.15, cdb.Amount)
	})

}
