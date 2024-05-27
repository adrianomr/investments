//go:generate mockgen -source cdi_repository.go -destination mock/cdi_repository_mock.go -package mock
package repositories

import (
	"context"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/database/sqlDB"
	"github.com/google/uuid"

	"github.com/adrianomr/investments/src/domain/models"
)

type CdiRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*models.Cdi, error)
	Create(ctx context.Context, cdi *models.Cdi) error
	FindAllOrderByCreatedAtDesc(ctx context.Context) ([]models.Cdi, error)
}

type CdiDBRepository struct{}

func NewCdiRepository() CdiRepository {
	return &CdiDBRepository{}
}

func (r *CdiDBRepository) Create(ctx context.Context, cdi *models.Cdi) error {
	query := "insert into cdi (id, rate, date) values ($1, $2, $3)"
	return sqlDB.NewStatement(ctx, query, cdi.ID, cdi.Rate, cdi.Date).Execute()
}

func (r *CdiDBRepository) FindById(ctx context.Context, id uuid.UUID) (*models.Cdi, error) {
	query := "select id, rate, date from cdi where id = $1"
	return sqlDB.NewQuery[models.Cdi](ctx, query, id).One()
}

func (r *CdiDBRepository) FindAllOrderByCreatedAtDesc(ctx context.Context) ([]models.Cdi, error) {
	query := "select id, rate, date from cdi order by created_at desc"
	return sqlDB.NewQuery[models.Cdi](ctx, query).Many()
}
