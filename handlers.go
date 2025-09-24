package main

import "github.com/gofiber/fiber/v2"

type URLHandler struct {
	urlService *URLService
	config     *Config
}

func NewURLHandler(cfg *Config) *URLHandler {
	return &URLHandler{
		urlService: NewURLService(),
		config:     cfg,
	}
}

// HealthCheck
func (h *URLHandler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": "url shortener",
		"version": "1.0.0",
	})
}

// ShortenURL
func (h *URLHandler) ShortenURL(c *fiber.Ctx) error {
	var req ShortenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if req.URL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL is required",
		})
	}

	urlModel, err := h.urlService.ShortenURL(req.URL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := ShortenResponse{
		OriginalURL: urlModel.OriginalURL,
		ShortCode:   urlModel.ShortCode,
		ShortURL:    h.config.BaseURL + "/" + urlModel.ShortCode,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
