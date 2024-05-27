package repositories

import (
	"context"
	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateAndFindCdi(t *testing.T) {
	repository := repositories.NewCdiRepository()

	t.Run("Should save and find cdi", func(t *testing.T) {
		datasets := []string{"clear-data.sql"}
		err := pc.Dataset(basePath, datasets...)
		assert.NoError(t, err)
		cdi := &models.Cdi{
			ID:   uuid.New(),
			Rate: 10.5,
			Date: time.Date(2024, 01, 01, 01, 01, 01, 00, time.UTC),
		}

		err = repository.Create(context.Background(), cdi)
		assert.NoError(t, err)

		result, err := repository.FindById(context.Background(), cdi.ID)
		assert.NoError(t, err)
		assert.Equal(t, cdi, result)
	})
}
