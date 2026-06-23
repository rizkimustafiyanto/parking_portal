package service

import (
	"fmt"
	"strings"

	dbmodel "backend/internal/common/model"
	"backend/internal/modules/user/dto"
	usermodel "backend/internal/modules/user/model"
	"backend/internal/modules/user/repository"
	pagedto "backend/pkg/dto"
	"backend/pkg/password"

	"github.com/google/uuid"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(req dto.CreateUserRequest) error {
	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	role := strings.TrimSpace(req.Role)
	if role == "" {
		role = "member"
	}

	user := usermodel.User{
		BaseModel: dbmodel.BaseModel{
			ID: uuid.New(),
		},
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     role,
		Balance:  0,
	}

	return s.repo.Create(&user)
}

func (s *service) GetByID(id string) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return toResponse(user), nil
}

func (s *service) GetAll(query pagedto.PaginationDTO, filter dto.ListUserRequest) ([]dto.UserResponse, int64, error) {
	query.Normalize()

	users, total, err := s.repo.FindAll(query, filter.Search, filter.Role)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for i := range users {
		responses = append(responses, *toResponse(&users[i]))
	}

	return responses, total, nil
}

func (s *service) Update(id string, req dto.UpdateUserRequest) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(req.Name) != "" {
		user.Name = req.Name
	}

	if strings.TrimSpace(req.Email) != "" {
		user.Email = req.Email
	}

	if strings.TrimSpace(req.Role) != "" {
		user.Role = req.Role
	}

	if strings.TrimSpace(req.Password) != "" {
		hashedPassword, err := password.Hash(req.Password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	return s.repo.Update(user)
}

func (s *service) TopUpBalance(id string, req dto.TopUpBalanceRequest) (*dto.UserResponse, error) {
	updated, err := s.repo.TopUpBalance(id, req.Amount)
	if err != nil {
		return nil, fmt.Errorf("top up balance: %w", err)
	}

	return toResponse(updated), nil
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}

func toResponse(user *usermodel.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}
