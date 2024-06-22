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

		cdb.Update(cdis_may, []models.CdbOrder{models.CdbOrder{
			UserID: "123",
			Type:   models.CdbOrderTypeBuy,
			Amount: 90940.38,
			Date:   time.Now(),
		}}, time.Date(2024, 05, 01, 12, 12, 12, 0, time.UTC))

		assert.Equal(t, 91589.46, cdb.Amount)
	})

	t.Run("Should update CDB several transactions occurs", func(t *testing.T) {

		cdb.Update(cdis_march, []models.CdbOrder{
			{
				UserID: "123",
				Type:   models.CdbOrderTypeBuy,
				Amount: 83798.91,
				Date:   createDate(2024, 03, 01),
			},
			{
				UserID: "123",
				Type:   models.CdbOrderTypeBuy,
				Amount: 1728,
				Date:   createDate(2024, 03, 10),
			},
			{
				UserID: "123",
				Type:   models.CdbOrderTypeBuy,
				Amount: 5553,
				Date:   createDate(2024, 03, 15),
			},
			{
				UserID: "123",
				Type:   models.CdbOrderTypeSell,
				Amount: -1486.25,
				Date:   createDate(2024, 03, 21),
			},
		}, time.Date(2024, 05, 01, 12, 12, 12, 0, time.UTC))

		assert.Equal(t, 90319.74, cdb.Amount)
	})

}

// R$ 83798,91 - R$ 84.010,09 = 207,17
// R$ 85.734,08 - R$ 85.842,04 = 107,96
// R$ 91.395,04 - R$ 91.510,13 = 115,09
// R$ 90.023,88 - R$ 90.241,06 = 217,18
