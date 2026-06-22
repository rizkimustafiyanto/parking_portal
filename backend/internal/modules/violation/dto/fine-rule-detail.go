package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateFineRuleDetailRequest struct {
	RuleVersionID uuid.UUID `json:"rule_version_id" validate:"required"`
	RuleType      string    `json:"rule_type"`
	Key           string    `json:"key" validate:"required"`
	Value         string    `json:"value" validate:"required"`
}

type UpdateFineRuleDetailRequest struct {
	RuleVersionID uuid.UUID `json:"rule_version_id"`
	RuleType      string    `json:"rule_type"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
}

type ListFineRuleDetailRequest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`

	Search string `form:"search"`
}

type FineRuleDetailResponse struct {
	ID            string    `json:"id"`
	RuleVersionID string    `json:"rule_version_id"`
	RuleType      string    `json:"rule_type"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type FineRuleDetailThrow1 struct {
	ID            string `json:"id"`
	RuleVersionID string `json:"rule_version_id"`
	RuleType      string `json:"rule_type"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}
