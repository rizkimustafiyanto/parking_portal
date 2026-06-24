package repository

import (
	"fmt"
	"strings"
	"time"

	dbmodel "backend/internal/common/model"
	invoiceconst "backend/internal/modules/invoice/constans"
	invoiceModel "backend/internal/modules/invoice/model"
	paymentconst "backend/internal/modules/payment-transaction/constans"
	"backend/internal/modules/payment-transaction/model"
	usermodel "backend/internal/modules/user/model"
	"backend/pkg/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(user *model.PaymentTransaction) error
	ProcessMemberBalancePayment(invoiceID string, internalTransactionID string, amount float64, status paymentconst.PaymentStatus, scenario paymentconst.PaymentScenario, paidAt time.Time) error
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

func (r *repository) ProcessMemberBalancePayment(invoiceID string, internalTransactionID string, amount float64, status paymentconst.PaymentStatus, scenario paymentconst.PaymentScenario, paidAt time.Time) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var inv invoiceModel.Invoice
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", invoiceID).First(&inv).Error; err != nil {
			return err
		}

		if inv.Status == invoiceconst.InvoicePaid {
			return fmt.Errorf("invoice is already paid")
		}

		var member usermodel.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", inv.MemberID).First(&member).Error; err != nil {
			return err
		}

		if amount > 0 && amount != inv.Amount {
			return fmt.Errorf("payment amount must match invoice amount")
		}

		if status == paymentconst.PaymentSuccess && member.Balance < inv.Amount {
			return fmt.Errorf("insufficient member balance")
		}

		if status == paymentconst.PaymentSuccess {
			member.Balance -= inv.Amount
			if err := tx.Model(&member).Update("balance", member.Balance).Error; err != nil {
				return err
			}
		}

		payment := model.PaymentTransaction{
			BaseModel: dbmodel.BaseModel{
				ID: uuid.New(),
			},
			InvoiceID:             inv.ID,
			InternalTransactionID: internalTransactionID,
			Amount:                inv.Amount,
			Status:                status,
			Scenario:              scenario,
			PaidAt:                paidAt,
		}

		if err := tx.Create(&payment).Error; err != nil {
			return err
		}

		if status == paymentconst.PaymentSuccess {
			if err := tx.Model(&invoiceModel.Invoice{}).Where("id = ?", invoiceID).Update("status", invoiceconst.InvoicePaid).Error; err != nil {
				return err
			}
		}

		return nil
	})
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
