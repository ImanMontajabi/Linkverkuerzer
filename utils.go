package main

import (
	"crypto/rand"
	"math/big"
	"net/url"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func IsValidURL(str string) bool {
	// Basic validation
	if str == "" {
		return false
	}

	// Parse URL
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	// Check scheme and host
	if u.Scheme == "" || u.Host == "" {
		return false
	}

	// Allow only HTTP and HTTPS
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

func GenerateShortCode(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func NormalizeURL(rawURL string) string {
	// Trim whitespace
	rawURL = strings.TrimSpace(rawURL)

	// Add scheme if missing
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	// Parse and reconstruct to normalize
	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL // Return original if parsing fails
	}

	// Convert host to lowercase
	u.Host = strings.ToLower(u.Host)

	// Remove trailing slash from path if it's just "/"
	if u.Path == "/" {
		u.Path = ""
	}

	return u.String()
}

func ValidateURLLength(rawURL string, maxLength int) bool {
	return len(rawURL) <= maxLength
}
