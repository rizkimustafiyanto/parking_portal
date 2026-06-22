package service

import (
	"backend/internal/modules/violation/dto"
	pagedto "backend/pkg/dto"
)

type Service interface {
	Create(req dto.CreateViolationRequest) error
	GetByID(id string) (*dto.ViolationResponse, error)
	GetAll(query pagedto.PaginationDTO, filter dto.ListViolationRequest) ([]dto.ViolationResponse, int64, error)
	Update(id string, req dto.UpdateViolationRequest) error
	Delete(id string) error
}
