package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/domain/usecases"
	"github.com/adrianomr/investments/src/infra/repositories/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewCdbOrderCreate(t *testing.T) {
	uc := usecases.NewCdbOrderCreate()
	assert.NotNil(t, uc)
}

func TestCreateCdbOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockCdbOrderRepository(ctrl)
	uc := usecases.CdbOrderCreate{Repo: repo}

	t.Run("should create order", func(t *testing.T) {
		ctx := context.Background()
		order := &models.CdbOrder{}
		repo.EXPECT().Create(ctx, order).Return(nil)

		result, err := uc.Execute(context.Background(), order)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, order, result)
	})

	t.Run("should fail when error on db", func(t *testing.T) {
		ctx := context.Background()
		order := &models.CdbOrder{}
		repo.EXPECT().Create(ctx, order).Return(errors.New("DB error"))

		_, err := uc.Execute(context.Background(), order)
		assert.Error(t, err)
	})

}
