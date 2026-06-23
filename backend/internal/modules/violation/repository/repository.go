package repository

import (
	"strings"
	"time"

	"backend/internal/modules/violation/model"
	pagedto "backend/pkg/dto"

	"gorm.io/gorm"
)

type Repository interface {
	Create(violation *model.Violation) error
	FindByID(id string) (*model.Violation, error)
	FindAll(query pagedto.PaginationDTO, search string) ([]model.Violation, int64, error)
	Update(violation *model.Violation) error
	Delete(id string) error
	FindActiveFineRuleVersion() (*model.FineRuleVersion, error)
	FindLatestFineRuleVersion() (*model.FineRuleVersion, error)
	CountUnpaidViolationsByPlateSince(plateNumber string, since time.Time) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(violation *model.Violation) error {
	return r.db.Create(violation).Error
}

func (r *repository) FindByID(id string) (*model.Violation, error) {
	var violation model.Violation

	err := r.db.
		Preload("Officer").
		Preload("FineRuleVersion.Publish").
		Preload("FineRuleVersion.Details").
		Where("id = ?", id).
		First(&violation).
		Error
	if err != nil {
		return nil, err
	}

	return &violation, nil
}

func (r *repository) FindAll(query pagedto.PaginationDTO, search string) ([]model.Violation, int64, error) {
	var violations []model.Violation
	var total int64

	db := r.db.Model(&model.Violation{})

	if trimmedSearch := strings.TrimSpace(search); trimmedSearch != "" {
		pattern := "%" + trimmedSearch + "%"
		db = db.Where("plate_number ILIKE ? OR location ILIKE ?", pattern, pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Preload("Officer").
		Preload("FineRuleVersion.Publish").
		Preload("FineRuleVersion.Details").
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&violations).
		Error

	return violations, total, err
}

func (r *repository) Update(violation *model.Violation) error {
	return r.db.Save(violation).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&model.Violation{}, "id = ?", id).Error
}

func (r *repository) FindActiveFineRuleVersion() (*model.FineRuleVersion, error) {
	var version model.FineRuleVersion

	err := r.db.
		Preload("Publish").
		Preload("Details").
		Where("is_active = ?", true).
		Order("version_number DESC").
		First(&version).
		Error
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func (r *repository) FindLatestFineRuleVersion() (*model.FineRuleVersion, error) {
	var version model.FineRuleVersion

	err := r.db.
		Preload("Publish").
		Preload("Details").
		Order("version_number DESC").
		First(&version).
		Error
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func (r *repository) CountUnpaidViolationsByPlateSince(plateNumber string, since time.Time) (int64, error) {
	var total int64

	err := r.db.
		Model(&model.Violation{}).
		Joins("JOIN invoices ON invoices.violation_id = violations.id").
		Where("violations.plate_number = ?", plateNumber).
		Where("violations.occurred_at >= ?", since).
		Where("invoices.status <> ?", "PAID").
		Count(&total).
		Error

	return total, err
}
