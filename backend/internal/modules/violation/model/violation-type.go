package model

import (
	common "backend/internal/common/model"
	usermodel "backend/internal/modules/user/model"

	"github.com/google/uuid"
)

type ViolationType struct {
	common.BaseModel
	Code        string    `gorm:"column:code;uniqueIndex;not null"`
	Name        string    `gorm:"column:name;not null"`
	BaseAmount  float64   `gorm:"column:base_amount;not null"`
	CreatedByID  uuid.UUID `gorm:"column:created_by_id"`

	CreatedBy usermodel.User `gorm:"foreignKey:CreatedByID;references:ID"`
}
