//go:generate mockgen -source cdb_repository.go -destination mock/cdb_repository_mock.go -package mock
package repositories

import (
	"context"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/google/uuid"

	"github.com/adrianomr/investments/src/domain/models"
)

type CdbRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*models.Cdb, error)
	Create(ctx context.Context, cdb *models.Cdb) error
}

type CdbDBRepository struct{}

func NewCdbRepository() CdbRepository {
	return &CdbDBRepository{}
}

func (r *CdbDBRepository) Create(ctx context.Context, cdb *models.Cdb) error {
	query := "insert into cdb (id, user_id, investment_type, percentage) values ($1, $2, $3, $4)"
	return sqlDB.NewStatement(ctx, query, cdb.ID, cdb.UserID, cdb.Type, cdb.Percentage).Execute()
}

func (r *CdbDBRepository) FindById(ctx context.Context, id uuid.UUID) (*models.Cdb, error) {
	query := "select id, user_id, amount, percentage, investment_type from cdb where id = $1"
	return sqlDB.NewQuery[models.Cdb](ctx, query, id).One()
}
