package database

import (
	"errors"
	"fmt"
	"strings"

	commonmodel "backend/internal/common/model"
	"backend/internal/config"
	"backend/internal/modules/user/model"
	"backend/pkg/password"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, cfg *config.Config) error {
	if !cfg.SeedDatabase {
		return nil
	}

	var existing model.User
	if err := db.Where("email = ?", cfg.SeedAdminEmail).First(&existing).Error; err == nil {
		return nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	role := strings.TrimSpace(cfg.SeedAdminRole)
	if role == "" {
		role = "admin"
	}

	hashedPassword, err := password.Hash(cfg.SeedAdminPassword)
	if err != nil {
		return fmt.Errorf("hash seed password: %w", err)
	}

	user := model.User{
		BaseModel: gormModelBase(),
		Name:      cfg.SeedAdminName,
		Email:     cfg.SeedAdminEmail,
		Password:  hashedPassword,
		Role:      role,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func gormModelBase() commonmodel.BaseModel {
	return commonmodel.BaseModel{
		ID: uuid.New(),
	}
}
