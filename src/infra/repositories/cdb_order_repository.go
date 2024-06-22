//go:generate mockgen -source cdb_order_repository.go -destination mock/cdb_order_repository_mock.go -package mock
package repositories

import (
	"context"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/google/uuid"

	"github.com/adrianomr/investments/src/domain/models"
)

type CdbOrderRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*models.CdbOrder, error)
	Create(ctx context.Context, cdbOrder *models.CdbOrder) error
	FindAllByCdbId(ctx context.Context, id uuid.UUID) ([]models.CdbOrder, error)
}

type CdbOrderDBRepository struct{}

func NewCdbOrderRepository() CdbOrderRepository {
	return &CdbOrderDBRepository{}
}

func (r *CdbOrderDBRepository) FindById(ctx context.Context, id uuid.UUID) (*models.CdbOrder, error) {
	query := "select id, user_id, type, amount, cdb_id, date from cdb_order where id = $1"
	return sqlDB.NewQuery[models.CdbOrder](ctx, query, id).One()
}

func (r *CdbOrderDBRepository) Create(ctx context.Context, cdbOrder *models.CdbOrder) error {
	query := "insert into cdb_order (id, user_id, type, amount, cdb_id, date) values ($1, $2, $3, $4, $5, $6)"
	return sqlDB.NewStatement(ctx, query, cdbOrder.ID, cdbOrder.UserID, cdbOrder.Type, cdbOrder.Amount, cdbOrder.CdbId, cdbOrder.Date).Execute()
}

func (r *CdbOrderDBRepository) FindAllByCdbId(ctx context.Context, id uuid.UUID) ([]models.CdbOrder, error) {
	query := "select id, user_id, type, amount, cdb_id, date from cdb_order where cdb_id = $1"
	return sqlDB.NewQuery[models.CdbOrder](ctx, query, id).Many()
}
