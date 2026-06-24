package service

import (
	"backend/internal/modules/payment-transaction/dto"
	pagedto "backend/pkg/dto"
)

type ChargeResult struct {
	Status        string
	TransactionID string
}

type PaymentService interface {
	Charge(invoiceID string, amount float64, scenario string) ChargeResult
}

type Service interface {
	Create(actorUserID string, actorRole string, req dto.CreatePaymentTransactionRequest) error

	GetByID(id string) (*dto.PaymentTransactionResponse, error)

	GetAll(query pagedto.PaginationDTO, filter dto.ListPaymentTransactionRequest) ([]dto.PaymentTransactionResponse, int64, error)

	Update(id string, req dto.UpdatePaymentTransactionRequest) error

	Delete(id string) error
}
