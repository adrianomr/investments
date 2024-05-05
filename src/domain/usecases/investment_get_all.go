//go:generate mockgen -source investment_get_all.go -destination mock/investment_get_all_mock.go -package mock
package usecases

import (
	"context"

	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
)

type IInvestmentGetAll interface {
	Execute(ctx context.Context) ([]models.Investment, error)
}

func NewInvestmentGetAll() IInvestmentGetAll {
	return &InvestmentGetAll{Repo: repositories.NewInvestmentMemoryRepository()}
}

type InvestmentGetAll struct {
	Repo repositories.InvestmentRepository
}

func (uc *InvestmentGetAll) Execute(ctx context.Context) ([]models.Investment, error) {
	return uc.Repo.FindAll(ctx)
}
