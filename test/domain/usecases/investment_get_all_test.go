package usecases

import (
	"context"
	"fmt"
	"testing"

	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/domain/usecases"
	"github.com/adrianomr/investments/src/infra/repositories/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewInvestmentGetAll(t *testing.T) {
	uc := usecases.NewInvestmentGetAll()
	assert.NotNil(t, uc)
}

func TestInvestmentGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock.NewMockInvestmentRepository(ctrl)
	uc := usecases.InvestmentGetAll{Repo: repo}

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		expected := []models.Investment{
			{ID: 1, Name: "1"},
			{ID: 2, Name: "2"},
		}
		repo.EXPECT().FindAll(ctx).Return(expected, nil)

		result, err := uc.Execute(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expected, result)
	})

	t.Run("fail", func(t *testing.T) {
		ctx := context.Background()
		repo.EXPECT().FindAll(ctx).Return(nil, fmt.Errorf("could not find"))

		result, err := uc.Execute(context.Background())
		assert.Error(t, err)
		assert.Empty(t, result)
	})
}
