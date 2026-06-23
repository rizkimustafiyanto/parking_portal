package repository

import (
	"strings"

	"backend/internal/modules/invoice/model"
	"backend/pkg/dto"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.Invoice) error
	FindByID(id string) (*model.Invoice, error)
	FindByMemberID(memberID string) ([]model.Invoice, error)
	FindAll(query dto.PaginationDTO, search string, status string) ([]model.Invoice, int64, error)
	Update(user *model.Invoice) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(invoice *model.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *repository) Update(invoice *model.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *repository) Delete(id string) error {

	return r.db.
		Delete(&model.Invoice{}, "id = ?", id).
		Error
}

func (r *repository) FindByID(id string) (*model.Invoice, error) {

	var invoice model.Invoice

	err := r.db.
		Preload("Member").
		Preload("Violation.Officer").
		Preload("Violation.FineRuleVersion.Publish").
		Preload("Violation.FineRuleVersion.Details").
		Preload("Payment").
		Where("id = ?", id).
		First(&invoice).
		Error

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (r *repository) FindByMemberID(memberID string) ([]model.Invoice, error) {
	var invoices []model.Invoice

	err := r.db.
		Preload("Member").
		Preload("Violation.Officer").
		Preload("Violation.FineRuleVersion.Publish").
		Preload("Violation.FineRuleVersion.Details").
		Preload("Payment").
		Where("member_id = ?", memberID).
		Order("created_at DESC").
		Find(&invoices).
		Error

	return invoices, err
}

func (r *repository) FindAll(query dto.PaginationDTO, search string, status string) ([]model.Invoice, int64, error) {
	var invoice []model.Invoice
	var total int64

	db := r.db.Model(&model.Invoice{})

	if trimmedStatus := strings.TrimSpace(status); trimmedStatus != "" {
		db = db.Where("status = ?", trimmedStatus)
	}

	if trimmedSearch := strings.TrimSpace(search); trimmedSearch != "" {
		pattern := "%" + trimmedSearch + "%"
		db = db.Where("CAST(amount AS TEXT) ILIKE ? OR CAST(status AS TEXT) ILIKE ?", pattern, pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Preload("Member").
		Preload("Violation.Officer").
		Preload("Violation.FineRuleVersion.Publish").
		Preload("Violation.FineRuleVersion.Details").
		Preload("Payment").
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&invoice).
		Error

	return invoice, total, err
}
