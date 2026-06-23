package repository

import (
	"strings"

	"backend/internal/modules/payment-transaction/model"
	"backend/pkg/dto"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.PaymentTransaction) error
	FindByID(id string) (*model.PaymentTransaction, error)
	FindAll(query dto.PaginationDTO, search string, status string, scenario string) ([]model.PaymentTransaction, int64, error)
	Update(user *model.PaymentTransaction) error
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

func (r *repository) Create(payments *model.PaymentTransaction) error {
	return r.db.Create(payments).Error
}

func (r *repository) Update(payments *model.PaymentTransaction) error {
	return r.db.Save(payments).Error
}

func (r *repository) Delete(id string) error {

	return r.db.
		Delete(&model.PaymentTransaction{}, "id = ?", id).
		Error
}

func (r *repository) FindByID(id string) (*model.PaymentTransaction, error) {

	var payments model.PaymentTransaction

	err := r.db.
		Where("id = ?", id).
		First(&payments).
		Error

	if err != nil {
		return nil, err
	}

	return &payments, nil
}

func (r *repository) FindAll(query dto.PaginationDTO, search string, status string, scenario string) ([]model.PaymentTransaction, int64, error) {
	var payments []model.PaymentTransaction
	var total int64

	db := r.db.Model(&model.PaymentTransaction{})

	if trimmedStatus := strings.TrimSpace(status); trimmedStatus != "" {
		db = db.Where("status = ?", trimmedStatus)
	}

	if trimmedScenario := strings.TrimSpace(scenario); trimmedScenario != "" {
		db = db.Where("scenario = ?", trimmedScenario)
	}

	if trimmedSearch := strings.TrimSpace(search); trimmedSearch != "" {
		pattern := "%" + trimmedSearch + "%"
		db = db.Where("CAST(invoice_id AS TEXT) ILIKE ? OR CAST(amount AS TEXT) ILIKE ? OR internal_transaction_id ILIKE ?", pattern, pattern, pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&payments).
		Error

	return payments, total, err
}
