package seed

import (
	"backend/internal/config"
	"backend/internal/modules/user/model"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, cfg *config.Config) error {
	if !cfg.SeedDatabase {
		return nil
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if cfg.SeedResetData {
			if err := resetSeedData(tx); err != nil {
				return err
			}
		}

		var existingCount int64
		if err := tx.Model(&model.User{}).Where("email = ?", cfg.SeedAdminEmail).Count(&existingCount).Error; err != nil {
			return err
		}
		if existingCount > 0 {
			return nil
		}

		admin, err := seedAdmin(tx, cfg)
		if err != nil {
			return err
		}

		if err := seedUserData(tx, admin); err != nil {
			return err
		}
		if err := seedViolationDomain(tx, admin); err != nil {
			return err
		}
		return seedFinanceDomain(tx)
	})
}
