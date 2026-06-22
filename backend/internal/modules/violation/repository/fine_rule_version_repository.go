package repository

import (
	"strings"

	"backend/internal/modules/violation/model"
	pagedto "backend/pkg/dto"

	"gorm.io/gorm"
)

type FineRuleVersionRepository interface {
	Create(version *model.FineRuleVersion) error
	FindByID(id string) (*model.FineRuleVersion, error)
	FindAll(query pagedto.PaginationDTO, search string) ([]model.FineRuleVersion, int64, error)
	Update(version *model.FineRuleVersion) error
	Delete(id string) error
}

type fineRuleVersionRepository struct {
	db *gorm.DB
}

func NewFineRuleVersionRepository(db *gorm.DB) FineRuleVersionRepository {
	return &fineRuleVersionRepository{db: db}
}

func (r *fineRuleVersionRepository) Create(version *model.FineRuleVersion) error {
	return r.db.Create(version).Error
}

func (r *fineRuleVersionRepository) FindByID(id string) (*model.FineRuleVersion, error) {
	var version model.FineRuleVersion

	err := r.db.
		Preload("Publish").
		Preload("Details").
		Where("id = ?", id).
		First(&version).
		Error
	if err != nil {
		return nil, err
	}

	return &version, nil
}

func (r *fineRuleVersionRepository) FindAll(query pagedto.PaginationDTO, search string) ([]model.FineRuleVersion, int64, error) {
	var versions []model.FineRuleVersion
	var total int64

	db := r.db.Model(&model.FineRuleVersion{})

	if trimmedSearch := strings.TrimSpace(search); trimmedSearch != "" {
		pattern := "%" + trimmedSearch + "%"
		db = db.Where("CAST(version_number AS TEXT) ILIKE ?", pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Preload("Publish").
		Preload("Details").
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&versions).
		Error

	return versions, total, err
}

func (r *fineRuleVersionRepository) Update(version *model.FineRuleVersion) error {
	return r.db.Save(version).Error
}

func (r *fineRuleVersionRepository) Delete(id string) error {
	return r.db.Delete(&model.FineRuleVersion{}, "id = ?", id).Error
}
