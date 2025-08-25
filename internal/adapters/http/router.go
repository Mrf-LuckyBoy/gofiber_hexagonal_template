package http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, bh *BookHandler) {
	api := app.Group("/api")
	books := api.Group("/books")
	books.Post("/", bh.Create)
	books.Get("/", bh.List)
	books.Get(":id", bh.Get)
	books.Put(":id", bh.Update)
	books.Delete(":id", bh.Delete)
}
