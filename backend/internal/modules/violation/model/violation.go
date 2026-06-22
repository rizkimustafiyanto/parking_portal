package model

import (
	common "backend/internal/common/model"
	usermodel "backend/internal/modules/user/model"
	"time"

	"github.com/google/uuid"
)

type Violation struct {
	common.BaseModel
	PlateNumber string    `gorm:"column:plate_number"`
	Location    string    `gorm:"column:location"`
	OccurredAt  time.Time `gorm:"column:occurred_at"`
	PhotoURL    string    `gorm:"column:photo_url"`
	OfficerID   uuid.UUID `gorm:"column:officer_id"`

	FineRuleVersionID uuid.UUID       `gorm:"column:fine_rule_version_id"`
	Officer           usermodel.User  `gorm:"foreignKey:OfficerID;references:ID"`
	FineRuleVersion   FineRuleVersion `gorm:"foreignKey:FineRuleVersionID;references:ID"`
}
