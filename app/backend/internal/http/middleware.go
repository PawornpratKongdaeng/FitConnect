package http

import (
	"github.com/gofiber/contrib/fiberjwt/v2"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(secret string) fiber.Handler {
	return fiberjwt.New(fiberjwt.Config{SigningKey: []byte(secret), ContextKey: "user"})
}
