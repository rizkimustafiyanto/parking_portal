package dto

import (
	"time"

	userdto "backend/internal/modules/user/dto"

	"github.com/google/uuid"
)

type CreateFineRuleVersionRequest struct {
	VersionNumber uint      `json:"version_number" validate:"required"`
	IsActive      bool      `json:"is_active"`
	PublishedBy   uuid.UUID `json:"published_by"`
}

type UpdateFineRuleVersionRequest struct {
	VersionNumber uint      `json:"version_number"`
	IsActive      bool      `json:"is_active"`
	PublishedBy   uuid.UUID `json:"published_by"`
}

type ListFineRuleVersionRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`

	Search string `form:"search"`
}

type FineRuleVersionResponse struct {
	ID            string                 `json:"id"`
	VersionNumber uint                   `json:"version_number"`
	IsActive      bool                   `json:"is_active"`
	Publish       userdto.UserThrow      `json:"publish"`
	Details       []FineRuleDetailThrow1 `json:"details"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

type FineRuleVersionThrow1 struct {
	ID            string `json:"id"`
	VersionNumber uint   `json:"version_number"`
	IsActive      bool   `json:"is_active"`
}
