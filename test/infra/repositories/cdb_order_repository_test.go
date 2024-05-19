package repositories

import (
	"context"
	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndFindCdbOrder(t *testing.T) {
	repository := repositories.NewCdbOrderRepository()

	t.Run("Should save and find cdb order", func(t *testing.T) {
		datasets := []string{"clear-data.sql", "insert-cdb.sql"}
		err := pc.Dataset(basePath, datasets...)
		assert.NoError(t, err)
		cdbOrder := &models.CdbOrder{
			ID:     uuid.New(),
			UserID: "1",
			Type:   models.CdbOrderTypeBuy,
			Amount: 150.55,
			CdbId:  uuid.MustParse("b0dae81c-e55c-4ca8-b635-2f3087d6b590"),
		}

		err = repository.Create(context.Background(), cdbOrder)
		assert.NoError(t, err)

		result, err := repository.FindById(context.Background(), cdbOrder.ID)
		assert.NoError(t, err)
		assert.Equal(t, cdbOrder, result)
	})
}
