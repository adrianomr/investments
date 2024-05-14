//go:generate mockgen -source cdb_order_repository.go -destination mock/cdb_order_repository_mock.go -package mock
package repositories

import (
	"context"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"

	"github.com/adrianomr/investments/src/domain/models"
)

type CdbOrderRepository interface {
	Create(ctx context.Context, cdbOrder *models.CdbOrder) (*models.CdbOrder, error)
}

type CdbOrderDBRepository struct{}

func NewCdbOrderRepository() CdbOrderRepository {
	return &CdbOrderDBRepository{}
}

func (r *CdbOrderDBRepository) Create(ctx context.Context, cdbOrder *models.CdbOrder) (*models.CdbOrder, error) {
	query := "insert into cdb_order (id, user_id, type, amount) values ($1, $2, $3, $4)"
	sqlDB.NewStatement(ctx, query, cdbOrder.ID, cdbOrder.UserID, cdbOrder.Type, cdbOrder.Amount).Execute()
	return cdbOrder, nil
}
