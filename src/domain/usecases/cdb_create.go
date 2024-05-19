//go:generate mockgen -source cdb_create.go -destination mock/cdb_create_mock.go -package mock
package usecases

import (
	"context"
	"github.com/google/uuid"

	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
)

type ICdbCreate interface {
	Execute(ctx context.Context, cdb *models.Cdb) (*models.Cdb, error)
}

func NewCdbCreate() ICdbCreate {
	return &CdbCreate{Repo: repositories.NewCdbRepository()}
}

type CdbCreate struct {
	Repo repositories.CdbRepository
}

func (uc *CdbCreate) Execute(ctx context.Context, cdb *models.Cdb) (*models.Cdb, error) {
	cdb.ID = uuid.New()
	return cdb, uc.Repo.Create(ctx, cdb)
}
