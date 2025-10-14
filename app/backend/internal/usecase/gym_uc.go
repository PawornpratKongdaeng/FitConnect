package usecase

import (
	"context"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"github.com/PawornpratKongdaeng/FitConnect/internal/repo"
)

type GymUC struct{ Repo *repo.GymRepo }

func (g *GymUC) Near(ctx context.Context, lat, lng, km float64) ([]domain.Gym, error) {
	return g.Repo.Near(ctx, lat, lng, km)
}
