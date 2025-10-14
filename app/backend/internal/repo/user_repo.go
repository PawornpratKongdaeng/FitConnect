package repo

import (
	"context"
	"time"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct{ C *mongo.Collection }

func NewUserRepo(db *mongo.Database) *UserRepo { return &UserRepo{C: db.Collection("users")} }

func (r *UserRepo) Create(ctx context.Context, u *domain.User) error {
	u.CreatedAt = time.Now().Unix()
	_, err := r.C.InsertOne(ctx, u)
	return err
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	err := r.C.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	return &u, err
}
