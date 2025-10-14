package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewServer(origin string) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{AllowOrigins: origin, AllowHeaders: "*", AllowCredentials: true, AllowMethods: "GET,POST,PUT,PATCH,DELETE"}))
	return app
}
