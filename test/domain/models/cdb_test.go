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
		Percentage: 0,
		Type:       models.CdbInvestmentTypeCdi,
	}

	t.Run("Should update CDB amount when no orders created", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{})

		assert.Equal(t, float64(0), cdb.Amount)
	})

	t.Run("Should update CDB amount when one order created", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{models.CdbOrder{
			UserID: "123",
			Type:   models.CdbOrderTypeBuy,
			Amount: 155.55,
		}})

		assert.Equal(t, 155.55, cdb.Amount)
	})

	t.Run("Should update CDB amount when first buy is today", func(t *testing.T) {
		cdb.Update([]models.Cdi{}, []models.CdbOrder{models.CdbOrder{
			UserID: "123",
			Type:   models.CdbOrderTypeBuy,
			Amount: 155.55,
			Date:   time.Now(),
		}})

		assert.Equal(t, 155.55, cdb.Amount)
	})

}
