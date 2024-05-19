package repositories

import (
	"context"
	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndFindCdb(t *testing.T) {
	repository := repositories.NewCdbRepository()

	t.Run("Should save and find cdb", func(t *testing.T) {
		datasets := []string{"clear-data.sql"}
		err := pc.Dataset(basePath, datasets...)
		assert.NoError(t, err)
		cdb := &models.Cdb{
			ID:             uuid.New(),
			UserID:         "1",
			Percentage:     0.5,
			InvestmentType: models.CdbInvestmentTypeCdi,
		}

		err = repository.Create(context.Background(), cdb)
		assert.NoError(t, err)

		result, err := repository.FindById(context.Background(), cdb.ID)
		assert.NoError(t, err)
		assert.Equal(t, cdb, result)
	})
}
