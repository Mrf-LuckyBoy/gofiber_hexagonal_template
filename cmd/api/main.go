package main

import (
	httpadapter "github.com/Mrf-LuckyBoy/test-go/internal/adapters/http"
	"github.com/Mrf-LuckyBoy/test-go/pkg/config"
	"github.com/Mrf-LuckyBoy/test-go/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Loading config from env
	cfg := config.Load()

	// putting env to use in the app
	c := BuildContainer(cfg)

	// starting fiber app
	app := fiber.New()

	// api path to check if the server is running
	// http://localhost:3000/health
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "ok"})
	})

	// register http adpter from dependency injection container
	httpadapter.RegisterRoutes(app, c.BookHandler, c.AuthHandler, c.UserHandler, cfg.JWTSecret)

	// port server start
	addr := ":" + cfg.Port
	logger.L.Println("starting server on", addr)
	if err := app.Listen(addr); err != nil {
		panic(err)
	}
}
