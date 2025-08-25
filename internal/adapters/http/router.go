package http

import (
	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, bookHandler *handlers.BookHandler, ah *handlers.AuthHandler, uh *handlers.UserHandler, secret string) {
	// , authHandler *handlers.AuthHandler, weatherHandler *handlers.WeatherHandler, secret string) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	books := v1.Group("/books")
	books.Post("/", bookHandler.Create)
	books.Get("/", bookHandler.List)
	books.Get(":id", bookHandler.Get)
	books.Put(":id", bookHandler.Update)
	books.Delete(":id", bookHandler.Delete)

	users := v1.Group("/users")
	users.Get("/", uh.List)
}
