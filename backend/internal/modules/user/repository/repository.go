package repository

import (
	"fmt"
	"strings"

	"backend/internal/modules/user/model"
	"backend/pkg/dto"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *model.User) error

	FindByID(id string) (*model.User, error)

	FindByEmail(email string) (*model.User, error)

	FindAll(query dto.PaginationDTO, search string, role string) ([]model.User, int64, error)

	Update(user *model.User) error

	TopUpBalance(id string, amount float64) (*model.User, error)

	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *repository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *repository) TopUpBalance(id string, amount float64) (*model.User, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	var updated model.User
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&user).Error; err != nil {
			return err
		}

		user.Balance += amount
		if err := tx.Model(&user).Update("balance", user.Balance).Error; err != nil {
			return err
		}

		updated = user
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r *repository) Delete(id string) error {

	return r.db.
		Delete(&model.User{}, "id = ?", id).
		Error
}

func (r *repository) FindByID(id string) (*model.User, error) {

	var user model.User

	err := r.db.
		Where("id = ?", id).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindByEmail(email string) (*model.User, error) {

	var user model.User

	err := r.db.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repository) FindAll(query dto.PaginationDTO, search string, role string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := r.db.Model(&model.User{})

	if trimmedRole := strings.TrimSpace(role); trimmedRole != "" {
		db = db.Where("role = ?", trimmedRole)
	}

	if trimmedSearch := strings.TrimSpace(search); trimmedSearch != "" {
		pattern := "%" + trimmedSearch + "%"
		db = db.Where("name ILIKE ? OR email ILIKE ?", pattern, pattern)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.
		Order(query.SortBy + " " + query.OrderBy).
		Offset(query.Offset()).
		Limit(query.Limit).
		Find(&users).
		Error

	return users, total, err
}
