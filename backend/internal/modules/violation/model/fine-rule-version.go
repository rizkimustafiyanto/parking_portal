package model

import (
	common "backend/internal/common/model"
	usermodel "backend/internal/modules/user/model"

	"github.com/google/uuid"
)

type FineRuleVersion struct {
	common.BaseModel
	VersionNumber uint `gorm:"column:version_number"`
	IsActive      bool
	PublishedBy   uuid.UUID `gorm:"column:published_by"`

	Publish usermodel.User   `gorm:"foreignKey:PublishedBy;references:ID"`
	Details []FineRuleDetail `gorm:"foreignKey:RuleVersionID;references:ID"`
}
