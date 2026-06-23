package dto

import (
	pStatus "backend/internal/modules/invoice/constans"
	paymentDto "backend/internal/modules/payment-transaction/dto"
	userDto "backend/internal/modules/user/dto"
	"time"

	"github.com/google/uuid"
)

type CreateInvoiceRequest struct {
	ViolationID uuid.UUID             `json:"violation_id"`
	MemberID    string                `json:"member_id"`
	Status      pStatus.InvoiceStatus `json:"status"`
}

type UpdateInvoiceRequest struct {
	ViolationID *uuid.UUID             `json:"violation_id"`
	MemberID    *string                `json:"member_id"`
	Status      *pStatus.InvoiceStatus `json:"status"`
}

type ListInvoiceRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`

	Search string `form:"search"`
	Status string `form:"status"`
}

type InvoiceResponse struct {
	ID         string                             `json:"id"`
	ViolationID uuid.UUID                          `json:"violation_id"`
	Amount     float64                            `json:"amount"`
	Status     pStatus.InvoiceStatus              `json:"status"`
	Member     userDto.UserThrow2                 `json:"member"`
	Violation  ViolationHistoryThrow               `json:"violation"`
	Payment    paymentDto.PaymentTransactionThrow  `json:"payment"`
	CreatedAt  time.Time                          `json:"created_at"`
	UpdatedAt  time.Time                          `json:"updated_at"`
}

type ViolationFineRuleDetailThrow struct {
	ID            string `json:"id"`
	RuleType      string `json:"rule_type"`
	Key           string `json:"key"`
	Value         string `json:"value"`
	RuleVersionID string `json:"rule_version_id"`
}

type ViolationFineRuleVersionThrow struct {
	ID            string                        `json:"id"`
	VersionNumber uint                          `json:"version_number"`
	IsActive      bool                          `json:"is_active"`
	Details       []ViolationFineRuleDetailThrow `json:"details"`
}

type ViolationHistoryThrow struct {
	ID              string                     `json:"id"`
	PlateNumber     string                     `json:"plate_number"`
	Location        string                     `json:"location"`
	OccurredAt      time.Time                  `json:"occurred_at"`
	PhotoURL        string                     `json:"photo_url"`
	FineRuleVersion ViolationFineRuleVersionThrow `json:"fine_rule_version"`
}

type InvoiceThrow struct {
	ID          string                `json:"id"`
	ViolationID uuid.UUID             `json:"violation_id"`
	Amount      float64               `json:"amount"`
	Status      pStatus.InvoiceStatus `json:"status"`
}
