package main

import (
	"errors"
	"time"
)

type URLService struct {
	repo   *URLRepository
	config *Config
}

func NewURLService(repo *URLRepository, config *Config) *URLService {
	return &URLService{
		repo:   repo,
		config: config,
	}
}

func (s *URLService) ShortenURL(originalURL string) (*URL, error) {
	// Normalize URL
	normalizedURL := NormalizeURL(originalURL)

	// Validate URL format
	if !IsValidURL(normalizedURL) {
		return nil, errors.New("invalid URL format")
	}

	// Validate URL length
	if !ValidateURLLength(normalizedURL, s.config.MaxURLLength) {
		return nil, errors.New("URL too long")
	}

	// Check if URL already exists
	existingURL, err := s.repo.GetByOriginalURL(normalizedURL)
	if err != nil {
		return nil, err
	}
	if existingURL != nil {
		return existingURL, nil
	}

	// Generate unique short code
	var shortCode string
	maxAttempts := 10
	for i := 0; i < maxAttempts; i++ {
		shortCode, err = GenerateShortCode(s.config.ShortCodeLen)
		if err != nil {
			return nil, errors.New("failed to generate short code")
		}

		// Check if short code already exists
		_, err = s.repo.GetByShortCode(shortCode)
		if err != nil && err.Error() == "URL not found" {
			// Short code is unique, break the loop
			break
		}
		if i == maxAttempts-1 {
			return nil, errors.New("failed to generate unique short code")
		}
	}

	// Create new URL record
	urlModel := &URL{
		OriginalURL: normalizedURL,
		ShortCode:   shortCode,
		ClickCount:  0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = s.repo.Create(urlModel)
	if err != nil {
		return nil, errors.New("failed to save URL")
	}

	return urlModel, nil
}

func (s *URLService) GetOriginalURL(shortCode string) (*URL, error) {
	if shortCode == "" {
		return nil, errors.New("short code is required")
	}

	url, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	// Increment click count
	if err := s.repo.UpdateClickCount(shortCode); err != nil {
		// Log error but don't fail the request
		// In production, you might want to use a proper logger here
	}

	return url, nil
}

func (s *URLService) GetStats(shortCode string) (*StatsResponse, error) {
	if shortCode == "" {
		return nil, errors.New("short code is required")
	}

	url, err := s.repo.GetByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	stats := &StatsResponse{
		ShortCode:   url.ShortCode,
		OriginalURL: url.OriginalURL,
		ClickCount:  url.ClickCount,
		CreatedAt:   url.CreatedAt,
	}

	return stats, nil
}

func (s *URLService) GetAllURLs() (*URLListResponse, error) {
	urls, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("failed to retrieve URLs")
	}

	count, err := s.repo.GetCount()
	if err != nil {
		return nil, errors.New("failed to get URL count")
	}

	response := &URLListResponse{
		URLs:  urls,
		Total: int(count),
	}

	return response, nil
}
