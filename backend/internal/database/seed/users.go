package seed

import (
	"fmt"
	"strings"

	"backend/internal/config"
	"backend/internal/modules/user/model"
	"backend/pkg/password"

	"gorm.io/gorm"
)

func seedAdmin(tx *gorm.DB, cfg *config.Config) (*model.User, error) {
	role := strings.TrimSpace(cfg.SeedAdminRole)
	if role == "" {
		role = "officer"
	}

	hashedPassword, err := password.Hash(cfg.SeedAdminPassword)
	if err != nil {
		return nil, fmt.Errorf("hash seed password: %w", err)
	}

	admin := &model.User{
		BaseModel: gormModelBase(),
		Name:      cfg.SeedAdminName,
		Email:     cfg.SeedAdminEmail,
		Password:  hashedPassword,
		Role:      role,
		Balance:   0,
	}

	if err := tx.Create(admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func createUser(tx *gorm.DB, name, email, rawPassword, role string, balance float64) (*model.User, error) {
	hashedPassword, err := password.Hash(rawPassword)
	if err != nil {
		return nil, fmt.Errorf("hash %s password: %w", role, err)
	}

	user := &model.User{
		BaseModel: gormModelBase(),
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      role,
		Balance:   balance,
	}

	if err := tx.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func seedUserData(tx *gorm.DB, admin *model.User) error {
	seedUsers := []struct {
		Name     string
		Email    string
		Password string
		Role     string
		Balance  float64
	}{
		{
			Name:     "Member One",
			Email:    "member1@example.com",
			Password: "member123",
			Role:     "member",
			Balance:  250000,
		},
		{
			Name:     "Member Two",
			Email:    "member2@example.com",
			Password: "member123",
			Role:     "member",
			Balance:  150000,
		},
		{
			Name:     "Officer Portal",
			Email:    "officer@example.com",
			Password: "officer123",
			Role:     "officer",
			Balance:  0,
		},
	}

	for i := range seedUsers {
		if _, err := createUser(
			tx,
			seedUsers[i].Name,
			seedUsers[i].Email,
			seedUsers[i].Password,
			seedUsers[i].Role,
			seedUsers[i].Balance,
		); err != nil {
			return err
		}
	}

	_ = admin
	return nil
}
