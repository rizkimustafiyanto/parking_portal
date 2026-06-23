package model

import (
	common "backend/internal/common/model"
	pStatus "backend/internal/modules/payment-transaction/constans"
	"time"

	"github.com/google/uuid"
)

type PaymentTransaction struct {
	common.BaseModel
	InvoiceID             uuid.UUID               `gorm:"column:invoice_id;not null"`
	InternalTransactionID string                  `gorm:"column:internal_transaction_id;not null"`
	Amount                float64                 `gorm:"not null;default:0"`
	Status                pStatus.PaymentStatus   `gorm:"type:text;not null"`
	Scenario              pStatus.PaymentScenario `gorm:"type:text;not null"`
	PaidAt                time.Time               `gorm:"column:paid_at;not null"`
}
