//go:generate mockgen -source investment_repository.go -destination mock/investment_repository_mock.go -package mock
package repositories

import (
	"context"

	"github.com/adrianomr/investments/src/domain/models"
)

var (
	memoryList = []models.Investment{
		{ID: 1, Name: "Investment 1"},
		{ID: 2, Name: "Investment 2"},
		{ID: 3, Name: "Investment 3"},
	}
)

type InvestmentRepository interface {
	FindAll(ctx context.Context) ([]models.Investment, error)
}

type InvestmentDBRepository struct{}

func NewInvestmentMemoryRepository() *InvestmentDBRepository {
	return &InvestmentDBRepository{}
}

func (r *InvestmentDBRepository) FindAll(_ context.Context) ([]models.Investment, error) {
	return memoryList, nil
}
