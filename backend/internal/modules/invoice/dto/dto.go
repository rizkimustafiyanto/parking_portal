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
	ID          string                             `json:"id"`
	ViolationID uuid.UUID                          `json:"violation_id"`
	Amount      float64                            `json:"amount"`
	Status      pStatus.InvoiceStatus              `json:"status"`
	Member      userDto.UserThrow2                 `json:"member"`
	Payment     paymentDto.PaymentTransactionThrow `json:"payment"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
}

type InvoiceThrow struct {
	ID          string                `json:"id"`
	ViolationID uuid.UUID             `json:"violation_id"`
	Amount      float64               `json:"amount"`
	Status      pStatus.InvoiceStatus `json:"status"`
}
