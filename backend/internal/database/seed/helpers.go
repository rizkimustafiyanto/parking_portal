package seed

import (
	"fmt"
	"strings"

	commonmodel "backend/internal/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func gormModelBase() commonmodel.BaseModel {
	return commonmodel.BaseModel{
		ID: uuid.New(),
	}
}

func resetSeedData(tx *gorm.DB) error {
	tables := []string{
		"payment_transactions",
		"invoices",
		"violations",
		"fine_rule_details",
		"fine_rule_versions",
		"violation_types",
		"users",
	}

	if err := tx.Exec("TRUNCATE TABLE " + strings.Join(tables, ", ") + " RESTART IDENTITY CASCADE").Error; err != nil {
		return fmt.Errorf("reset seed data: %w", err)
	}

	return nil
}
