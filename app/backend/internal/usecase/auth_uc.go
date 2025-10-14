package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/PawornpratKongdaeng/FitConnect/internal/domain"
	"github.com/PawornpratKongdaeng/FitConnect/internal/repo"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUC struct {
	Users  *repo.UserRepo
	Secret []byte
}

type RegisterInput struct{ Name, Email, Password string }

type LoginInput struct{ Email, Password string }

func (a *AuthUC) Register(ctx context.Context, in RegisterInput) (*domain.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), 12)
	u := &domain.User{Name: in.Name, Email: in.Email, Password: string(hash), Role: domain.RoleUser}
	if err := a.Users.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (a *AuthUC) Login(ctx context.Context, in LoginInput) (string, error) {
	u, err := a.Users.FindByEmail(ctx, in.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password)) != nil {
		return "", errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": u.ID, "role": string(u.Role), "exp": time.Now().Add(24 * time.Hour).Unix()})
	return token.SignedString(a.Secret)
}
