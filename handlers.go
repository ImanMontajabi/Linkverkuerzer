package main

import (
	"github.com/gofiber/fiber/v2"
)

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
