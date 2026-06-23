package model

import (
	common "backend/internal/common/model"
	pStatus "backend/internal/modules/invoice/constans"
	violationModel "backend/internal/modules/violation/model"
	paymentModel "backend/internal/modules/payment-transaction/model"
	userModel "backend/internal/modules/user/model"

	"github.com/google/uuid"
)

type Invoice struct {
	common.BaseModel
	ViolationID uuid.UUID             `gorm:"column:violation_id;not null"`
	MemberID    uuid.UUID             `gorm:"column:member_id;not null"`
	Amount      float64               `gorm:"not null;default:0"`
	Status      pStatus.InvoiceStatus `gorm:"type:text;not null"`

	Member   userModel.User                  `gorm:"foreignKey:MemberID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Violation violationModel.Violation       `gorm:"foreignKey:ViolationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Payment  paymentModel.PaymentTransaction `gorm:"foreignKey:InvoiceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
