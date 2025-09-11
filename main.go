package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"error":   err.Error(),
				"success": false,
			})
		},
		AppName: "Linkverkuerzer v1.0.0",
	})
	// Middleware

}
