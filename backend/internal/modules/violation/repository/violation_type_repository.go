package repository

import (
	"strings"

	"backend/internal/modules/violation/model"
	pagedto "backend/pkg/dto"

	"gorm.io/gorm"
)

type ViolationTypeRepository interface {
	Create(item *model.ViolationType) error
	FindByID(id string) (*model.ViolationType, error)
	FindByCode(code string) (*model.ViolationType, error)
	FindAll(query pagedto.PaginationDTO, search string) ([]model.ViolationType, int64, error)
	Update(item *model.ViolationType) error
	Delete(id string) error
}

type violationTypeRepository struct {
	db *gorm.DB
}

func NewViolationTypeRepository(db *gorm.DB) ViolationTypeRepository {
	return &violationTypeRepository{db: db}
}

func (r *violationTypeRepository) Create(item *model.ViolationType) error {
	return r.db.Create(item).Error
}

func (r *violationTypeRepository) FindByID(id string) (*model.ViolationType, error) {
	var item model.ViolationType
	if err := r.db.Preload("CreatedBy").Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *violationTypeRepository) FindByCode(code string) (*model.ViolationType, error) {
	var item model.ViolationType
	if err := r.db.Preload("CreatedBy").Where("code = ?", code).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *violationTypeRepository) FindAll(query pagedto.PaginationDTO, search string) ([]model.ViolationType, int64, error) {
	var items []model.ViolationType
	var total int64

	db := r.db.Model(&model.ViolationType{})
	if trimmed := strings.TrimSpace(search); trimmed != "" {
		pattern := "%" + trimmed + "%"
		db = db.Where("code ILIKE ? OR name ILIKE ?", pattern, pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Preload("CreatedBy").
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&items).Error

	return items, total, err
}

func (r *violationTypeRepository) Update(item *model.ViolationType) error {
	return r.db.Save(item).Error
}

func (r *violationTypeRepository) Delete(id string) error {
	return r.db.Delete(&model.ViolationType{}, "id = ?", id).Error
}
