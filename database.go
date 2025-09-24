package main

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func ConnectDatabase(config *Config) {
	var err error

	// Connect to SQLite
	DB, err = gorm.Open(sqlite.Open(config.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")
}

func MigrateDatabase() {
	err := DB.AutoMigrate(&URL{})
	if err != nil {
		log.Fatal("Failed to migrate Database:", err)
	}
	log.Println("Database migrated successfully")
}

func CloseDatabase() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting database instance: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}

// Repository methods
func (r *URLRepository) Create(url *URL) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) GetByShortCode(shortCode string) (*URL, error) {
	var url URL
	err := r.db.Where("short_code = ?", shortCode).First(&url).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("URL not found")
		}
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) GetByOriginalURL(originalURL string) (*URL, error) {
	var url URL
	err := r.db.Where("original_url = ?", originalURL).First(&url).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Not an error, just not found
		}
		return nil, err
	}
	return &url, nil
}

func (r *URLRepository) UpdateClickCount(shortCode string) error {
	return r.db.Model(&URL{}).Where("short_code = ?", shortCode).Update("click_count", gorm.Expr("click_count + ?", 1)).Error
}

func (r *URLRepository) GetAll() ([]URL, error) {
	var urls []URL
	err := r.db.Order("created_at desc").Find(&urls).Error
	return urls, err
}

func (r *URLRepository) GetCount() (int64, error) {
	var count int64
	err := r.db.Model(&URL{}).Count(&count).Error
	return count, err
}
