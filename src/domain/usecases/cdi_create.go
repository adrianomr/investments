//go:generate mockgen -source cdi_create.go -destination mock/cdi_create_mock.go -package mock
package usecases

import (
	"context"
	"github.com/google/uuid"

	"github.com/adrianomr/investments/src/domain/models"
	"github.com/adrianomr/investments/src/infra/repositories"
)

type ICdiCreate interface {
	Execute(ctx context.Context, cdi *models.Cdi) (*models.Cdi, error)
}

func NewCdiCreate() ICdiCreate {
	return &CdiCreate{Repo: repositories.NewCdiRepository()}
}

type CdiCreate struct {
	Repo repositories.CdiRepository
}

func (uc *CdiCreate) Execute(ctx context.Context, cdi *models.Cdi) (*models.Cdi, error) {
	cdi.ID = uuid.New()
	return cdi, uc.Repo.Create(ctx, cdi)
}
