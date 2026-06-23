package dto

import (
	"time"

	userdto "backend/internal/modules/user/dto"

	"github.com/google/uuid"
)

type CreateViolationRequest struct {
	PlateNumber       string    `json:"plate_number" validate:"required"`
	ViolationTypeCode string    `json:"violation_type_code" validate:"required"`
	Location          string    `json:"location" validate:"required"`
	OccurredAt        time.Time `json:"occurred_at" validate:"required"`
	PhotoURL          string    `json:"photo_url"`
	OfficerID         uuid.UUID `json:"officer_id" validate:"required"`
	FineRuleVersionID uuid.UUID `json:"fine_rule_version_id"`
}

type UpdateViolationRequest struct {
	PlateNumber       string    `json:"plate_number"`
	ViolationTypeCode string    `json:"violation_type_code"`
	Location          string    `json:"location"`
	OccurredAt        time.Time `json:"occurred_at"`
	PhotoURL          string    `json:"photo_url"`
	OfficerID         uuid.UUID `json:"officer_id"`
	FineRuleVersionID uuid.UUID `json:"fine_rule_version_id"`
}

type ListViolationRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`

	Search string `form:"search"`
}

type ViolationOfficerThrow struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FineRuleVersionThrow struct {
	ID            string `json:"id"`
	VersionNumber uint   `json:"version_number"`
	IsActive      bool   `json:"is_active"`
}

type ViolationResponse struct {
	ID              string               `json:"id"`
	PlateNumber     string               `json:"plate_number"`
	ViolationTypeCode string             `json:"violation_type_code"`
	Location        string               `json:"location"`
	OccurredAt      time.Time            `json:"occurred_at"`
	PhotoURL        string               `json:"photo_url"`
	Officer         userdto.UserThrow    `json:"officer"`
	FineRuleVersion FineRuleVersionThrow `json:"fine_rule_version"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
}
