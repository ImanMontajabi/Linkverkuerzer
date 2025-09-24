package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load config
	config := LoadConfig()

	// Connect to database
	ConnectDatabase(config)
	MigrateDatabase()

	// Initialize repository
	urlRepo := NewURLRepository(DB)

	// Initialize service
	urlService := NewURLService(urlRepo, config)

	// Initialize handler
	urlHandler := NewURLHandler(urlService, config)

	// Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(ErrorResponse{
				Error:   err.Error(),
				Success: false,
			})
		},
		AppName: "Linkverkuerzer v1.0.0",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} | ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Routes (order is important)
	app.Get("/", urlHandler.HealthCheck)
	app.Post("/shorten", urlHandler.ShortenUrl)
	app.Get("/stats/:code", urlHandler.GetStats)
	app.Get("/urls", urlHandler.GetAllUrls)
	app.Get("/:code", urlHandler.RedirectUrl)

	// Start server
	log.Printf("Server starting on port %s", config.Port)
	log.Fatal(app.Listen(":" + config.Port))
}
