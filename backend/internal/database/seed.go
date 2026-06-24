package database

import (
	"backend/internal/config"
	backendseed "backend/internal/database/seed"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cfg *config.Config) error {
	return backendseed.Run(db, cfg)
}
