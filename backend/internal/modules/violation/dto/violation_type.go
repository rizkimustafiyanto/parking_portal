package dto

import (
	"time"

	userdto "backend/internal/modules/user/dto"

	"github.com/google/uuid"
)

type CreateViolationTypeRequest struct {
	Code       string    `json:"code" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	BaseAmount float64   `json:"base_amount" validate:"required"`
	CreatedByID uuid.UUID `json:"created_by_id"`
}

type UpdateViolationTypeRequest struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	BaseAmount float64 `json:"base_amount"`
	CreatedByID uuid.UUID `json:"created_by_id"`
}

type ListViolationTypeRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`

	Search string `form:"search"`
}

type ViolationTypeResponse struct {
	ID         string            `json:"id"`
	Code       string            `json:"code"`
	Name       string            `json:"name"`
	BaseAmount float64           `json:"base_amount"`
	CreatedBy  userdto.UserThrow `json:"created_by"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}
