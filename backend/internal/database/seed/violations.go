package seed

import (
	"time"

	"backend/internal/modules/user/model"
	violationmodel "backend/internal/modules/violation/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func seedViolationDomain(tx *gorm.DB, admin *model.User) error {
	ruleVersion, err := createFineRuleVersion(tx, admin.ID, 1, true)
	if err != nil {
		return err
	}

	if err := createFineRuleDetails(tx, ruleVersion.ID); err != nil {
		return err
	}

	violationTypes, err := createViolationTypes(tx, admin.ID)
	if err != nil {
		return err
	}

	officer, err := findSeedUserByEmail(tx, "officer@example.com")
	if err != nil {
		return err
	}

	return createViolations(tx, admin.ID, officer.ID, ruleVersion.ID, violationTypes)
}

func findSeedUserByEmail(tx *gorm.DB, email string) (*model.User, error) {
	var user model.User
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func createFineRuleVersion(tx *gorm.DB, publishedBy uuid.UUID, version uint, active bool) (*violationmodel.FineRuleVersion, error) {
	ruleVersion := &violationmodel.FineRuleVersion{
		BaseModel:     gormModelBase(),
		VersionNumber: version,
		IsActive:      active,
		PublishedBy:   publishedBy,
	}

	if err := tx.Create(ruleVersion).Error; err != nil {
		return nil, err
	}

	return ruleVersion, nil
}

func createFineRuleDetails(tx *gorm.DB, ruleVersionID uuid.UUID) error {
	details := []violationmodel.FineRuleDetail{
		{
			BaseModel:     gormModelBase(),
			RuleVersionID: ruleVersionID,
			RuleType:      "percentage",
			Key:           "late_payment_penalty",
			Value:         "10",
		},
		{
			BaseModel:     gormModelBase(),
			RuleVersionID: ruleVersionID,
			RuleType:      "fixed",
			Key:           "administration_fee",
			Value:         "25000",
		},
		{
			BaseModel:     gormModelBase(),
			RuleVersionID: ruleVersionID,
			RuleType:      "fixed",
			Key:           "collection_fee",
			Value:         "15000",
		},
	}

	for i := range details {
		if err := tx.Create(&details[i]).Error; err != nil {
			return err
		}
	}

	return nil
}

func createViolationTypes(tx *gorm.DB, adminID uuid.UUID) ([]violationmodel.ViolationType, error) {
	violationTypes := []violationmodel.ViolationType{
		{
			BaseModel:   gormModelBase(),
			Code:        "helmet",
			Name:        "Tidak memakai helm",
			BaseAmount:  50000,
			CreatedByID: adminID,
		},
		{
			BaseModel:   gormModelBase(),
			Code:        "wrong_lane",
			Name:        "Melawan arus",
			BaseAmount:  100000,
			CreatedByID: adminID,
		},
		{
			BaseModel:   gormModelBase(),
			Code:        "red_light",
			Name:        "Menerobos lampu merah",
			BaseAmount:  150000,
			CreatedByID: adminID,
		},
	}

	for i := range violationTypes {
		if err := tx.Create(&violationTypes[i]).Error; err != nil {
			return nil, err
		}
	}

	return violationTypes, nil
}

func createViolations(tx *gorm.DB, adminID, officerID, ruleVersionID uuid.UUID, violationTypes []violationmodel.ViolationType) error {
	violations := []violationmodel.Violation{
		{
			BaseModel:         gormModelBase(),
			PlateNumber:       "B 1234 CD",
			ViolationTypeCode: violationTypes[0].Code,
			Location:          "Jakarta Pusat",
			OccurredAt:        time.Date(2026, time.June, 24, 9, 30, 0, 0, time.UTC),
			PhotoURL:          "https://example.com/uploads/violation-sample-1.jpg",
			OfficerID:         adminID,
			FineRuleVersionID: ruleVersionID,
		},
		{
			BaseModel:         gormModelBase(),
			PlateNumber:       "B 4321 EF",
			ViolationTypeCode: violationTypes[1].Code,
			Location:          "Jakarta Selatan",
			OccurredAt:        time.Date(2026, time.June, 24, 11, 15, 0, 0, time.UTC),
			PhotoURL:          "https://example.com/uploads/violation-sample-2.jpg",
			OfficerID:         officerID,
			FineRuleVersionID: ruleVersionID,
		},
		{
			BaseModel:         gormModelBase(),
			PlateNumber:       "B 7788 GH",
			ViolationTypeCode: violationTypes[2].Code,
			Location:          "Jakarta Barat",
			OccurredAt:        time.Date(2026, time.June, 23, 17, 45, 0, 0, time.UTC),
			PhotoURL:          "https://example.com/uploads/violation-sample-3.jpg",
			OfficerID:         officerID,
			FineRuleVersionID: ruleVersionID,
		},
	}

	for i := range violations {
		if err := tx.Create(&violations[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
