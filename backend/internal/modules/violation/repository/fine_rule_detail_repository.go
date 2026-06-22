package repository

import (
	"strings"

	"backend/internal/modules/violation/model"
	pagedto "backend/pkg/dto"

	"gorm.io/gorm"
)

type FineRuleDetailRepository interface {
	Create(detail *model.FineRuleDetail) error
	FindByID(id string) (*model.FineRuleDetail, error)
	FindAll(query pagedto.PaginationDTO, search string) ([]model.FineRuleDetail, int64, error)
	Update(detail *model.FineRuleDetail) error
	Delete(id string) error
}

type fineRuleDetailRepository struct {
	db *gorm.DB
}

func NewFineRuleDetailRepository(db *gorm.DB) FineRuleDetailRepository {
	return &fineRuleDetailRepository{db: db}
}

func (r *fineRuleDetailRepository) Create(detail *model.FineRuleDetail) error {
	return r.db.Create(detail).Error
}

func (r *fineRuleDetailRepository) FindByID(id string) (*model.FineRuleDetail, error) {
	var detail model.FineRuleDetail

	err := r.db.Where("id = ?", id).First(&detail).Error
	if err != nil {
		return nil, err
	}

	return &detail, nil
}

func (r *fineRuleDetailRepository) FindAll(query pagedto.PaginationDTO, search string) ([]model.FineRuleDetail, int64, error) {
	var details []model.FineRuleDetail
	var total int64

	db := r.db.Model(&model.FineRuleDetail{})

	if trimmedSearch := strings.TrimSpace(search); trimmedSearch != "" {
		pattern := "%" + trimmedSearch + "%"
		db = db.Where(`"key" ILIKE ? OR "value" ILIKE ?`, pattern, pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&details).
		Error

	return details, total, err
}

func (r *fineRuleDetailRepository) Update(detail *model.FineRuleDetail) error {
	return r.db.Save(detail).Error
}

func (r *fineRuleDetailRepository) Delete(id string) error {
	return r.db.Delete(&model.FineRuleDetail{}, "id = ?", id).Error
}
