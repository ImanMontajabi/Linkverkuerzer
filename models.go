package main

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	OriginalURL string         `json:"original_url" gorm:"not null"`
	ShortCode   string         `json:"short_code" gorm:"unique;not null;index"`
	ClickCount  int            `json:"click_count" gorm:"default:0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ShortenRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type ShortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	ShortCode   string `json:"short_code"`
}

type StatsResponse struct {
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	ClickCount  int       `json:"click_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type URLListResponse struct {
	URLs  []URL `json:"urls"`
	Total int   `json:"total"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}
