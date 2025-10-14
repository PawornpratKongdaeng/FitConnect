package repo

import (
	"context"
	"time"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrainerRepo struct{ C *mongo.Collection }

func NewTrainerRepo(db *mongo.Database) *TrainerRepo {
	return &TrainerRepo{C: db.Collection("trainers")}
}

func (r *TrainerRepo) Create(ctx context.Context, t *domain.Trainer) error {
	t.CreatedAt = time.Now().Unix()
	_, err := r.C.InsertOne(ctx, t)
	return err
}

func (r *TrainerRepo) ListByRegion(ctx context.Context, region string, limit int64) ([]domain.Trainer, error) {
	cur, err := r.C.Find(ctx, bson.M{"region_code": region}, &options.FindOptions{Limit: &limit})
	if err != nil {
		return nil, err
	}
	var out []domain.Trainer
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}
