package usecase

import (
	"context"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"github.com/PawornpratKongdaeng/FitConnect/internal/repo"
)

type TrainerUC struct{ Repo *repo.TrainerRepo }

func (t *TrainerUC) Create(ctx context.Context, in domain.Trainer) error {
	return t.Repo.Create(ctx, &in)
}
func (t *TrainerUC) ListByRegion(ctx context.Context, region string, limit int64) ([]domain.Trainer, error) {
	return t.Repo.ListByRegion(ctx, region, limit)
}
