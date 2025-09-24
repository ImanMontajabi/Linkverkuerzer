package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type URLHandler struct {
	urlService *URLService
	config     *Config
}

func NewURLHandler(urlService *URLService, cfg *Config) *URLHandler {
	return &URLHandler{
		urlService: urlService,
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

// ShortenUrl - Create a shortened URL
func (h *URLHandler) ShortenUrl(c *fiber.Ctx) error {
	var req ShortenRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Invalid request body",
			Success: false,
		})
	}

	if req.URL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "URL is required",
			Success: false,
		})
	}

	urlModel, err := h.urlService.ShortenURL(req.URL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
		})
	}

	response := ShortenResponse{
		OriginalURL: urlModel.OriginalURL,
		ShortCode:   urlModel.ShortCode,
		ShortURL:    h.config.BaseURL + "/" + urlModel.ShortCode,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// RedirectUrl - Redirect to original URL
func (h *URLHandler) RedirectUrl(c *fiber.Ctx) error {
	shortCode := c.Params("code")

	if shortCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Short code is required",
			Success: false,
		})
	}

	urlModel, err := h.urlService.GetOriginalURL(shortCode)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error:   "URL not found",
				Success: false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error:   "Internal server error",
			Success: false,
		})
	}

	return c.Redirect(urlModel.OriginalURL, fiber.StatusMovedPermanently)
}

// GetStats - Get statistics for a shortened URL
func (h *URLHandler) GetStats(c *fiber.Ctx) error {
	shortCode := c.Params("code")

	if shortCode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   "Short code is required",
			Success: false,
		})
	}

	stats, err := h.urlService.GetStats(shortCode)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error:   "URL not found",
				Success: false,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error:   "Internal server error",
			Success: false,
		})
	}

	return c.JSON(stats)
}

// GetAllUrls - Get all shortened URLs
func (h *URLHandler) GetAllUrls(c *fiber.Ctx) error {
	urlList, err := h.urlService.GetAllURLs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error:   "Failed to retrieve URLs",
			Success: false,
		})
	}

	return c.JSON(urlList)
}
