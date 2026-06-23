package service

import (
	"fmt"
	dbmodel "backend/internal/common/model"
	userdto "backend/internal/modules/user/dto"
	"backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/model"
	"backend/internal/modules/violation/repository"
	pagedto "backend/pkg/dto"

	"github.com/google/uuid"
)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(req dto.CreateViolationRequest) error {
	fineRuleVersionID := req.FineRuleVersionID
	if fineRuleVersionID == uuid.Nil {
		activeVersion, err := s.resolveActiveFineRuleVersion()
		if err != nil {
			return err
		}
		fineRuleVersionID = activeVersion.ID
	}

	violation := model.Violation{
		BaseModel: dbmodel.BaseModel{
			ID: uuid.New(),
		},
		PlateNumber:       req.PlateNumber,
		Location:          req.Location,
		OccurredAt:        req.OccurredAt,
		PhotoURL:          req.PhotoURL,
		OfficerID:         req.OfficerID,
		FineRuleVersionID: fineRuleVersionID,
	}

	return s.repo.Create(&violation)
}

func (s *service) GetByID(id string) (*dto.ViolationResponse, error) {
	violation, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return toResponse(violation), nil
}

func (s *service) GetAll(query pagedto.PaginationDTO, filter dto.ListViolationRequest) ([]dto.ViolationResponse, int64, error) {
	query.Normalize()

	violations, total, err := s.repo.FindAll(query, filter.Search)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.ViolationResponse, 0, len(violations))
	for i := range violations {
		responses = append(responses, *toResponse(&violations[i]))
	}

	return responses, total, nil
}

func (s *service) Update(id string, req dto.UpdateViolationRequest) error {
	violation, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if req.PlateNumber != "" {
		violation.PlateNumber = req.PlateNumber
	}
	if req.Location != "" {
		violation.Location = req.Location
	}
	if !req.OccurredAt.IsZero() {
		violation.OccurredAt = req.OccurredAt
	}
	if req.PhotoURL != "" {
		violation.PhotoURL = req.PhotoURL
	}
	if req.OfficerID != uuid.Nil {
		violation.OfficerID = req.OfficerID
	}
	if req.FineRuleVersionID != uuid.Nil {
		violation.FineRuleVersionID = req.FineRuleVersionID
	}

	return s.repo.Update(violation)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}

func toResponse(violation *model.Violation) *dto.ViolationResponse {
	return &dto.ViolationResponse{
		ID:          violation.ID.String(),
		PlateNumber: violation.PlateNumber,
		Location:    violation.Location,
		OccurredAt:  violation.OccurredAt,
		PhotoURL:    violation.PhotoURL,
		Officer: userdto.UserThrow{
			ID:    violation.Officer.ID.String(),
			Name:  violation.Officer.Name,
			Email: violation.Officer.Email,
		},
		FineRuleVersion: dto.FineRuleVersionThrow{
			ID:            violation.FineRuleVersion.ID.String(),
			VersionNumber: violation.FineRuleVersion.VersionNumber,
			IsActive:      violation.FineRuleVersion.IsActive,
		},
		CreatedAt: violation.CreatedAt,
		UpdatedAt: violation.UpdatedAt,
	}
}

func (s *service) resolveActiveFineRuleVersion() (*model.FineRuleVersion, error) {
	version, err := s.repo.FindActiveFineRuleVersion()
	if err == nil {
		return version, nil
	}

	latest, latestErr := s.repo.FindLatestFineRuleVersion()
	if latestErr == nil {
		return latest, nil
	}

	return nil, fmt.Errorf("no active fine rule version available")
}
