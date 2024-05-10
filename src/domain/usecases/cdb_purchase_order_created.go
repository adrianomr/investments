//go:generate mockgen -source investment_get_all.go -destination mock/investment_get_all_mock.go -package mock
package usecases

import (
	"context"
	"github.com/google/uuid"

	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
)

type ICdbOrderCreate interface {
	Execute(ctx context.Context, order *models.CdbOrder) (*models.CdbOrder, error)
}

func NewCdbOrderCreate() ICdbOrderCreate {
	return &CdbOrderCreate{Repo: repositories.NewInvestmentMemoryRepository()}
}

type CdbOrderCreate struct {
	Repo repositories.CdbOrderRepository
}

func (uc *CdbOrderCreate) Execute(ctx context.Context, order *models.CdbOrder) (*models.CdbOrder, error) {
	order.ID = uuid.New()
	return uc.Repo.Create(ctx, order)
}
