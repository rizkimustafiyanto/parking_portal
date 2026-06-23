package service

import (
	"backend/internal/modules/invoice/dto"
	pagedto "backend/pkg/dto"
)

type Service interface {
	Create(req dto.CreateInvoiceRequest) error

	GetByID(id string) (*dto.InvoiceResponse, error)

	GetAll(query pagedto.PaginationDTO, filter dto.ListInvoiceRequest) ([]dto.InvoiceResponse, int64, error)

	Update(id string, req dto.UpdateInvoiceRequest) error

	Delete(id string) error
}
