package model

import (
	common "backend/internal/common/model"
)

type User struct {
	common.BaseModel
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Balance  float64 `gorm:"not null;default:0"`
}

func (u User) GetID() string {
	return u.ID.String()
}

func (u User) GetRole() string {
	return u.Role
}
