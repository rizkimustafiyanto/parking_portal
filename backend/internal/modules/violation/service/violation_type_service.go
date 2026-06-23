package service

import (
	dbmodel "backend/internal/common/model"
	userdto "backend/internal/modules/user/dto"
	vdto "backend/internal/modules/violation/dto"
	"backend/internal/modules/violation/model"
	"backend/internal/modules/violation/repository"
	pagedto "backend/pkg/dto"

	"github.com/google/uuid"
)

type ViolationTypeService interface {
	Create(req vdto.CreateViolationTypeRequest) error
	GetByID(id string) (*vdto.ViolationTypeResponse, error)
	GetAll(query pagedto.PaginationDTO, filter vdto.ListViolationTypeRequest) ([]vdto.ViolationTypeResponse, int64, error)
	Update(id string, req vdto.UpdateViolationTypeRequest) error
	Delete(id string) error
}

type violationTypeService struct {
	repo repository.ViolationTypeRepository
}

func NewViolationTypeService(repo repository.ViolationTypeRepository) ViolationTypeService {
	return &violationTypeService{repo: repo}
}

func (s *violationTypeService) Create(req vdto.CreateViolationTypeRequest) error {
	item := model.ViolationType{
		BaseModel:   dbmodel.BaseModel{ID: uuid.New()},
		Code:        req.Code,
		Name:        req.Name,
		BaseAmount:  req.BaseAmount,
		CreatedByID: req.CreatedByID,
	}
	return s.repo.Create(&item)
}

func (s *violationTypeService) GetByID(id string) (*vdto.ViolationTypeResponse, error) {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return toViolationTypeResponse(item), nil
}

func (s *violationTypeService) GetAll(query pagedto.PaginationDTO, filter vdto.ListViolationTypeRequest) ([]vdto.ViolationTypeResponse, int64, error) {
	query.Normalize()
	items, total, err := s.repo.FindAll(query, filter.Search)
	if err != nil {
		return nil, 0, err
	}
	responses := make([]vdto.ViolationTypeResponse, 0, len(items))
	for i := range items {
		responses = append(responses, *toViolationTypeResponse(&items[i]))
	}
	return responses, total, nil
}

func (s *violationTypeService) Update(id string, req vdto.UpdateViolationTypeRequest) error {
	item, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if req.Code != "" {
		item.Code = req.Code
	}
	if req.Name != "" {
		item.Name = req.Name
	}
	if req.BaseAmount > 0 {
		item.BaseAmount = req.BaseAmount
	}
	if req.CreatedByID != uuid.Nil {
		item.CreatedByID = req.CreatedByID
	}
	return s.repo.Update(item)
}

func (s *violationTypeService) Delete(id string) error {
	return s.repo.Delete(id)
}

func toViolationTypeResponse(item *model.ViolationType) *vdto.ViolationTypeResponse {
	return &vdto.ViolationTypeResponse{
		ID:         item.ID.String(),
		Code:       item.Code,
		Name:       item.Name,
		BaseAmount: item.BaseAmount,
		CreatedBy: userdto.UserThrow{
			ID:    item.CreatedBy.ID.String(),
			Name:  item.CreatedBy.Name,
			Email: item.CreatedBy.Email,
		},
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
