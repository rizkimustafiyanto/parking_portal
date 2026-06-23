package dto

import (
	pStatus "backend/internal/modules/payment-transaction/constans"
	"time"

	"github.com/google/uuid"
)

type CreatePaymentTransactionRequest struct {
	InvoiceID             uuid.UUID               `json:"invoice_id"`
	InternalTransactionID string                  `json:"internal_transaction_id"`
	Amount                float64                 `json:"amount" validate:"min=0"`
	Status                pStatus.PaymentStatus   `json:"status"`
	Scenario              pStatus.PaymentScenario `json:"scenario"`
	PaidAt                time.Time               `json:"paid_at"`
}

type UpdatePaymentTransactionRequest struct {
	InternalTransactionID *string                  `json:"internal_transaction_id"`
	Amount                *float64                 `json:"amount"`
	Status                *pStatus.PaymentStatus   `json:"status"`
	Scenario              *pStatus.PaymentScenario `json:"scenario"`
	PaidAt                *time.Time               `json:"paid_at"`
}

type ListPaymentTransactionRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`

	Search   string `form:"search"`
	Status   string `form:"status"`
	Scenario string `form:"scenario"`
}

type PaymentTransactionResponse struct {
	ID                    string                  `json:"id"`
	InvoiceID             uuid.UUID               `json:"invoice_id"`
	InternalTransactionID string                  `json:"internal_transaction_id"`
	Amount                float64                 `json:"amount"`
	Status                pStatus.PaymentStatus   `json:"status"`
	Scenario              pStatus.PaymentScenario `json:"scenario"`
	PaidAt                time.Time               `json:"paid_at"`
	CreatedAt             time.Time               `json:"created_at"`
	UpdatedAt             time.Time               `json:"updated_at"`
}

type PaymentTransactionThrow struct {
	ID                    string                  `json:"id"`
	InvoiceID             uuid.UUID               `json:"invoice_id"`
	InternalTransactionID string                  `json:"internal_transaction_id"`
	Amount                float64                 `json:"amount"`
	Status                pStatus.PaymentStatus   `json:"status"`
	Scenario              pStatus.PaymentScenario `json:"scenario"`
	PaidAt                time.Time               `json:"paid_at"`
}
