package service

import (
	dbmodel "backend/internal/common/model"
	"backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/model"
	"backend/internal/modules/violation/repository"
	pagedto "backend/pkg/dto"

	"github.com/google/uuid"
)

type FineRuleDetailService interface {
	Create(req dto.CreateFineRuleDetailRequest) error
	GetByID(id string) (*dto.FineRuleDetailResponse, error)
	GetAll(query pagedto.PaginationDTO, filter dto.ListFineRuleDetailRequest) ([]dto.FineRuleDetailResponse, int64, error)
	Update(id string, req dto.UpdateFineRuleDetailRequest) error
	Delete(id string) error
}

type fineRuleDetailService struct {
	repo repository.FineRuleDetailRepository
}

func NewFineRuleDetailService(repo repository.FineRuleDetailRepository) FineRuleDetailService {
	return &fineRuleDetailService{repo: repo}
}

func (s *fineRuleDetailService) Create(req dto.CreateFineRuleDetailRequest) error {
	detail := model.FineRuleDetail{
		BaseModel:     dbmodel.BaseModel{ID: uuid.New()},
		RuleVersionID: req.RuleVersionID,
		RuleType:      req.RuleType,
		Key:           req.Key,
		Value:         req.Value,
	}

	return s.repo.Create(&detail)
}

func (s *fineRuleDetailService) GetByID(id string) (*dto.FineRuleDetailResponse, error) {
	detail, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return toFineRuleDetailResponse(detail), nil
}

func (s *fineRuleDetailService) GetAll(query pagedto.PaginationDTO, filter dto.ListFineRuleDetailRequest) ([]dto.FineRuleDetailResponse, int64, error) {
	query.Normalize()
	details, total, err := s.repo.FindAll(query, filter.Search)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.FineRuleDetailResponse, 0, len(details))
	for i := range details {
		responses = append(responses, *toFineRuleDetailResponse(&details[i]))
	}

	return responses, total, nil
}

func (s *fineRuleDetailService) Update(id string, req dto.UpdateFineRuleDetailRequest) error {
	detail, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if req.RuleVersionID != uuid.Nil {
		detail.RuleVersionID = req.RuleVersionID
	}
	if req.RuleType != "" {
		detail.RuleType = req.RuleType
	}
	if req.Key != "" {
		detail.Key = req.Key
	}
	if req.Value != "" {
		detail.Value = req.Value
	}
	return s.repo.Update(detail)
}

func (s *fineRuleDetailService) Delete(id string) error {
	return s.repo.Delete(id)
}

func toFineRuleDetailResponse(detail *model.FineRuleDetail) *dto.FineRuleDetailResponse {
	return &dto.FineRuleDetailResponse{
		ID:            detail.ID.String(),
		RuleVersionID: detail.RuleVersionID.String(),
		RuleType:      detail.RuleType,
		Key:           detail.Key,
		Value:         detail.Value,
		CreatedAt:     detail.CreatedAt,
		UpdatedAt:     detail.UpdatedAt,
	}
}
