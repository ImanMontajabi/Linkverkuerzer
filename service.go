package main

import "errors"

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

func (s *URLService) ShortenURL(originalURL string) (*URL, error) {
	if !IsValidURL(originalURL) {
		return nil, errors.New("invalid URL format")
	}
}
