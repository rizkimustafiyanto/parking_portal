package repository

import (
	"strings"

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
		Preload("FineRuleVersion").
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
		Preload("FineRuleVersion").
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
