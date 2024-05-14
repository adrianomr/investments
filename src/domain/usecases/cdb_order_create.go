//go:generate mockgen -source cdb_order_create.go -destination mock/cdb_order_create_mock.go -package mock
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
	return &CdbOrderCreate{Repo: repositories.NewCdbOrderRepository()}
}

type CdbOrderCreate struct {
	Repo repositories.CdbOrderRepository
}

func (uc *CdbOrderCreate) Execute(ctx context.Context, order *models.CdbOrder) (*models.CdbOrder, error) {
	order.ID = uuid.New()
	return uc.Repo.Create(ctx, order)
}
