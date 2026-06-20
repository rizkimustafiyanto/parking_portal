package service

import "backend/internal/modules/user/dto"
import pagedto "backend/pkg/dto"

type Service interface {
	Create(req dto.CreateUserRequest) error

	GetByID(id string) (*dto.UserResponse, error)

	GetAll(query pagedto.PaginationDTO, filter dto.ListUserRequest) ([]dto.UserResponse, int64, error)

	Update(id string, req dto.UpdateUserRequest) error

	Delete(id string) error
}
