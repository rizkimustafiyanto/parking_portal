package model

import (
	common "backend/internal/common/model"

	"github.com/google/uuid"
)

type FineRuleDetail struct {
	common.BaseModel
	RuleVersionID uuid.UUID `gorm:"column:rule_version_id"`
	RuleType      string    `gorm:"column:rule_type"`
	Key           string    `gorm:"column:key"`
	Value         string    `gorm:"column:value"`
}
