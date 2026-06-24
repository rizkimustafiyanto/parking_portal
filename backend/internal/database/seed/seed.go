package seed

import (
	"errors"

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

		var existing model.User
		if err := tx.Where("email = ?", cfg.SeedAdminEmail).First(&existing).Error; err == nil {
			return nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
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
