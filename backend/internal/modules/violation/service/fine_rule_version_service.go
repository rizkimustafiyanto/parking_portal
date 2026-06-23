package service

import (
	"errors"
	"fmt"

	dbmodel "backend/internal/common/model"
	userdto "backend/internal/modules/user/dto"
	"backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/model"
	"backend/internal/modules/violation/repository"
	pagedto "backend/pkg/dto"

	"github.com/google/uuid"
)

type FineRuleVersionService interface {
	Create(req dto.CreateFineRuleVersionRequest) error
	GetByID(id string) (*dto.FineRuleVersionResponse, error)
	GetAll(query pagedto.PaginationDTO, filter dto.ListFineRuleVersionRequest) ([]dto.FineRuleVersionResponse, int64, error)
	Update(id string, req dto.UpdateFineRuleVersionRequest) error
	Delete(id string) error
}

type fineRuleVersionService struct {
	repo repository.FineRuleVersionRepository
}

func NewFineRuleVersionService(repo repository.FineRuleVersionRepository) FineRuleVersionService {
	return &fineRuleVersionService{repo: repo}
}

func (s *fineRuleVersionService) Create(req dto.CreateFineRuleVersionRequest) error {
	version := model.FineRuleVersion{
		BaseModel:     dbmodel.BaseModel{ID: uuid.New()},
		VersionNumber: req.VersionNumber,
		IsActive:      req.IsActive,
		PublishedBy:   req.PublishedBy,
	}

	return s.repo.Create(&version)
}

func (s *fineRuleVersionService) GetByID(id string) (*dto.FineRuleVersionResponse, error) {
	version, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return toFineRuleVersionResponse(version), nil
}

func (s *fineRuleVersionService) GetAll(query pagedto.PaginationDTO, filter dto.ListFineRuleVersionRequest) ([]dto.FineRuleVersionResponse, int64, error) {
	query.Normalize()
	versions, total, err := s.repo.FindAll(query, filter.Search)
	if err != nil {
		return nil, 0, err
	}
	responses := make([]dto.FineRuleVersionResponse, 0, len(versions))
	for i := range versions {
		responses = append(responses, *toFineRuleVersionResponse(&versions[i]))
	}
	return responses, total, nil
}

func (s *fineRuleVersionService) Update(id string, req dto.UpdateFineRuleVersionRequest) error {
	return fmt.Errorf("fine rule version is immutable; create a new version instead")
}

func (s *fineRuleVersionService) Delete(id string) error {
	return errors.New("fine rule version is immutable and cannot be deleted")
}

func toFineRuleVersionResponse(version *model.FineRuleVersion) *dto.FineRuleVersionResponse {
	details := make([]dto.FineRuleDetailThrow1, 0, len(version.Details))
	for i := range version.Details {
		details = append(details, dto.FineRuleDetailThrow1{
			ID:       version.Details[i].ID.String(),
			Key:      version.Details[i].Key,
			Value:    version.Details[i].Value,
			RuleType: version.Details[i].RuleType,
		})
	}

	return &dto.FineRuleVersionResponse{
		ID:            version.ID.String(),
		VersionNumber: version.VersionNumber,
		IsActive:      version.IsActive,
		Publish: userdto.UserThrow{
			ID:    version.Publish.ID.String(),
			Name:  version.Publish.Name,
			Email: version.Publish.Email,
		},
		Details:   details,
		CreatedAt: version.CreatedAt,
		UpdatedAt: version.UpdatedAt,
	}
}
