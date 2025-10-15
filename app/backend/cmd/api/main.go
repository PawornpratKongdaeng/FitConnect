package main

import (
	"log"

	"github.com/PawornpratKongdaeng/FitConnect/internal/config"
	"github.com/PawornpratKongdaeng/FitConnect/internal/db"
	"github.com/PawornpratKongdaeng/FitConnect/internal/handler"
	ih "github.com/PawornpratKongdaeng/FitConnect/internal/http"
	"github.com/PawornpratKongdaeng/FitConnect/internal/repo"
	"github.com/PawornpratKongdaeng/FitConnect/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	cfg := config.Load()
	client, err := db.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	var database *mongo.Database = client.Database("fitconnect")

	app := ih.NewServer(cfg.ClientOrigin)

	users := repo.NewUserRepo(database)
	trainers := repo.NewTrainerRepo(database)
	gyms := repo.NewGymRepo(database)

	authH := &handler.AuthHandler{UC: &usecase.AuthUC{Users: users, Secret: []byte(cfg.JWTSecret)}}
	trainerH := &handler.TrainerHandler{UC: &usecase.TrainerUC{Repo: trainers}}
	gymH := &handler.GymHandler{UC: &usecase.GymUC{Repo: gyms}}

	api := app.Group("/api")
	api.Post("/auth/register", authH.Register)
	api.Post("/auth/login", authH.Login)
	api.Post("/trainers", trainerH.Create)
	api.Get("/trainers", trainerH.ListByRegion)
	api.Get("/gyms/near", gymH.Near)

	log.Fatal(app.Listen(":" + cfg.Port))

}
