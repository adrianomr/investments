//go:generate mockgen -source cdb_repository.go -destination mock/cdb_repository_mock.go -package mock
package repositories

import (
	"context"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"

	"github.com/adrianomr/investments/src/domain/models"
)

type CdbRepository interface {
	Create(ctx context.Context, cdb *models.Cdb) (*models.Cdb, error)
}

type CdbDBRepository struct{}

func NewCdbRepository() CdbRepository {
	return &CdbDBRepository{}
}

func (r *CdbDBRepository) Create(ctx context.Context, cdb *models.Cdb) (*models.Cdb, error) {
	query := "insert into cdb_order (id, user_id, investment_type, cdi_percentage, amount) values ($1, $2, $3, $4)"
	sqlDB.NewStatement(ctx, query, cdb.ID, cdb.UserID, cdb.InvestmentType, cdb.CdiPercentage, cdb.Amount).Execute()
	return cdb, nil
}
