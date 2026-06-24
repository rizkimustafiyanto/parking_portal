package database

import (
	"time"

	commonmodel "backend/internal/common/model"
	usermodel "backend/internal/modules/user/model"
	violationmodel "backend/internal/modules/violation/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func seedViolationData(db *gorm.DB, admin *usermodel.User) error {
	if err := seedFineRuleVersion(db, admin); err != nil {
		return err
	}

	return seedViolation(db, admin)
}

func seedFineRuleVersion(db *gorm.DB, admin *usermodel.User) error {
	var existingCount int64
	if err := db.Model(&violationmodel.FineRuleVersion{}).Where("version_number = ?", 1).Count(&existingCount).Error; err != nil {
		return err
	}
	if existingCount > 0 {
		return nil
	}

	version := violationmodel.FineRuleVersion{
		BaseModel: commonmodel.BaseModel{
			ID: uuid.New(),
		},
		VersionNumber: 1,
		IsActive:      true,
		PublishedBy:   admin.ID,
	}

	if err := db.Create(&version).Error; err != nil {
		return err
	}

	detail := violationmodel.FineRuleDetail{
		BaseModel: commonmodel.BaseModel{
			ID: uuid.New(),
		},
		RuleVersionID: version.ID,
		RuleType:      "speeding",
		Key:           "fine_amount",
		Value:         "500000",
	}

	return db.Create(&detail).Error
}

func seedViolation(db *gorm.DB, admin *usermodel.User) error {
	var existingCount int64
	if err := db.Model(&violationmodel.Violation{}).Where("plate_number = ?", "B 1234 CD").Count(&existingCount).Error; err != nil {
		return err
	}
	if existingCount > 0 {
		return nil
	}

	var versionCount int64
	if err := db.Model(&violationmodel.FineRuleVersion{}).Where("version_number = ?", 1).Count(&versionCount).Error; err != nil {
		return err
	}
	if versionCount == 0 {
		return nil
	}

	var version violationmodel.FineRuleVersion
	if err := db.Where("version_number = ?", 1).First(&version).Error; err != nil {
		return err
	}

	v := violationmodel.Violation{
		BaseModel: commonmodel.BaseModel{
			ID: uuid.New(),
		},
		PlateNumber:       "B 1234 CD",
		Location:          "Jakarta Pusat",
		OccurredAt:        time.Now().UTC(),
		PhotoURL:          "https://example.com/violation.jpg",
		OfficerID:         admin.ID,
		FineRuleVersionID: version.ID,
	}

	return db.Create(&v).Error
}
