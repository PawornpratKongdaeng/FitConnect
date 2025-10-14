package repo

import (
	"context"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GymRepo struct{ C *mongo.Collection }

func NewGymRepo(db *mongo.Database) *GymRepo { return &GymRepo{C: db.Collection("gyms")} }

func (r *GymRepo) Near(ctx context.Context, lat, lng, km float64) ([]domain.Gym, error) {
	m := km * 1000
	filter := bson.M{"location": bson.M{"$near": bson.M{"$geometry": bson.M{"type": "Point", "coordinates": []float64{lng, lat}}, "$maxDistance": m}}}
	cur, err := r.C.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var out []domain.Gym
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}
