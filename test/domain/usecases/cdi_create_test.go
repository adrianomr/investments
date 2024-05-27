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

func TestNewCdiCreate(t *testing.T) {
	uc := usecases.NewCdiCreate()
	assert.NotNil(t, uc)
}

func TestCreateCdi(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockCdiRepository(ctrl)
	uc := usecases.CdiCreate{Repo: repo}

	t.Run("should create cbd", func(t *testing.T) {
		ctx := context.Background()
		cdi := &models.Cdi{}
		repo.EXPECT().Create(ctx, cdi).Return(nil)

		result, err := uc.Execute(context.Background(), cdi)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, cdi, result)
	})

	t.Run("should fail when error on db", func(t *testing.T) {
		ctx := context.Background()
		cdi := &models.Cdi{}
		repo.EXPECT().Create(ctx, cdi).Return(errors.New("DB error"))

		_, err := uc.Execute(context.Background(), cdi)
		assert.Error(t, err)
	})

}
