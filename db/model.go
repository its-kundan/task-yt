package db

import "time"

type Video struct {
	ID            uint      `gorm:"primaryKey"`
	VideoID       string    `gorm:"uniqueIndex"`
	Title         string
	Description   string
	PublishedAt   time.Time `gorm:"index"`
	ThumbnailsURL string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func MigrateModels() {
	DB.AutoMigrate(&Video{})
}
