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

func TestNewCdbCreate(t *testing.T) {
	uc := usecases.NewCdbCreate()
	assert.NotNil(t, uc)
}

func TestCreateCdb(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockCdbRepository(ctrl)
	uc := usecases.CdbCreate{Repo: repo}

	t.Run("should create cbd", func(t *testing.T) {
		ctx := context.Background()
		cdb := &models.Cdb{}
		repo.EXPECT().Create(ctx, cdb).Return(nil)

		result, err := uc.Execute(context.Background(), cdb)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, cdb, result)
	})

	t.Run("should fail when error on db", func(t *testing.T) {
		ctx := context.Background()
		cdb := &models.Cdb{}
		repo.EXPECT().Create(ctx, cdb).Return(errors.New("DB error"))

		_, err := uc.Execute(context.Background(), cdb)
		assert.Error(t, err)
	})

}
