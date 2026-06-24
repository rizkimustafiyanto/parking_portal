package seed

import (
	"fmt"

	invoiceconst "backend/internal/modules/invoice/constans"
	invoicemodel "backend/internal/modules/invoice/model"
	paymentconst "backend/internal/modules/payment-transaction/constans"
	paymentmodel "backend/internal/modules/payment-transaction/model"
	violationmodel "backend/internal/modules/violation/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func seedFinanceDomain(tx *gorm.DB) error {
	memberOne, err := findSeedUserByEmail(tx, "member1@example.com")
	if err != nil {
		return err
	}

	memberTwo, err := findSeedUserByEmail(tx, "member2@example.com")
	if err != nil {
		return err
	}

	var violations []violationmodel.Violation
	if err := tx.Order("created_at asc").Find(&violations).Error; err != nil {
		return err
	}

	invoices, err := createInvoices(tx, memberOne.ID, memberTwo.ID, violations)
	if err != nil {
		return err
	}

	return createPaymentTransactions(tx, invoices)
}

func createInvoices(tx *gorm.DB, memberOneID, memberTwoID uuid.UUID, violations []violationmodel.Violation) ([]invoicemodel.Invoice, error) {
	if len(violations) < 3 {
		return nil, fmt.Errorf("seed finance requires at least 3 violations, got %d", len(violations))
	}

	invoices := []invoicemodel.Invoice{
		{
			BaseModel:   gormModelBase(),
			ViolationID: violations[0].ID,
			MemberID:    memberOneID,
			Amount:      50000,
			Status:      invoiceconst.InvoicePaid,
		},
		{
			BaseModel:   gormModelBase(),
			ViolationID: violations[1].ID,
			MemberID:    memberOneID,
			Amount:      100000,
			Status:      invoiceconst.InvoicePending,
		},
		{
			BaseModel:   gormModelBase(),
			ViolationID: violations[2].ID,
			MemberID:    memberTwoID,
			Amount:      150000,
			Status:      invoiceconst.InvoiceOverdue,
		},
	}

	for i := range invoices {
		if err := tx.Create(&invoices[i]).Error; err != nil {
			return nil, err
		}
	}

	return invoices, nil
}

func createPaymentTransactions(tx *gorm.DB, invoices []invoicemodel.Invoice) error {
	payments := []paymentmodel.PaymentTransaction{
		{
			BaseModel:             gormModelBase(),
			InvoiceID:             invoices[0].ID,
			InternalTransactionID: "TRX-SUCCESS-SEED-0001",
			Amount:                50000,
			Status:                paymentconst.PaymentSuccess,
			Scenario:              paymentconst.ScenarioSuccess,
			PaidAt:                timeDateUTC(2026, 6, 24, 10, 0),
		},
		{
			BaseModel:             gormModelBase(),
			InvoiceID:             invoices[1].ID,
			InternalTransactionID: "TRX-FAILED-SEED-0002",
			Amount:                100000,
			Status:                paymentconst.PaymentFailed,
			Scenario:              paymentconst.ScenarioFailure,
			PaidAt:                timeDateUTC(2026, 6, 24, 13, 0),
		},
		{
			BaseModel:             gormModelBase(),
			InvoiceID:             invoices[2].ID,
			InternalTransactionID: "TRX-TIMEOUT-SEED-0003",
			Amount:                150000,
			Status:                paymentconst.PaymentFailed,
			Scenario:              paymentconst.ScenarioTimeout,
			PaidAt:                timeDateUTC(2026, 6, 25, 8, 0),
		},
	}

	for i := range payments {
		if err := tx.Create(&payments[i]).Error; err != nil {
			return err
		}
	}

	return nil
}
