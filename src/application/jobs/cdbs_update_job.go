package jobs

import (
	"context"
	"github.com/adrianomr/investments/src/infra/repositories"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type CdbsUpdate struct {
	Repository         repositories.CdbRepository
	CdiRepository      repositories.CdiRepository
	CdbOrderRepository repositories.CdbOrderRepository
}

func (u CdbsUpdate) Execute() {
	ctx := context.Background()
	cdbs, err := u.Repository.FindCdbsToUpdate(ctx)
	if err != nil {
		log.Errorf("Failed to find cdbs to update: %v", err)
	}
	cdis, err := u.CdiRepository.FindAllOrderByCreatedAtDesc(ctx)
	if err != nil {
		log.Errorf("Failed to find cdis: %v", err)
	}
	for _, cdb := range cdbs {
		orders, err := u.CdbOrderRepository.FindAllByCdbId(ctx, cdb.ID)
		if err != nil {
			log.Errorf("Failed to find cdis: %v", err)
		}
		cdb.Update(cdis, orders, time.Now())
	}
}
func (u CdbsUpdate) ExecuteAfter() time.Duration {
	return 5 * time.Second
}

func NewCdbsUpdateJob() Job {
	return &CdbsUpdate{}
}
