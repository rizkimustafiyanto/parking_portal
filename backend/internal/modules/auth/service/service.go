package service

import (
	"errors"
	"fmt"

	"backend/internal/modules/auth/dto"
	"backend/internal/modules/user/repository"
	"backend/pkg/jwt"
	pwd "backend/pkg/password"
)

var ErrInvalidCredential = errors.New("invalid credential")

type Service interface {
	Login(req dto.LoginRequest) (string, error)
}

type service struct {
	userRepo repository.Repository
	secret   string
}

func NewService(userRepo repository.Repository, secret string) Service {
	return &service{userRepo: userRepo, secret: secret}
}

func (s *service) Login(req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if err := pwd.Compare(user.Password, req.Password); err != nil {
		return "", ErrInvalidCredential
	}

	token, err := jwt.Generate(user, s.secret)
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return token, nil
}
