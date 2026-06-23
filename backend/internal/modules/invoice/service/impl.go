package service

import (
	"fmt"
	"strings"

	dbmodel "backend/internal/common/model"
	"backend/internal/modules/invoice/dto"
	invoiceModel "backend/internal/modules/invoice/model"
	"backend/internal/modules/invoice/repository"
	paymentdto "backend/internal/modules/payment-transaction/dto"
	paymentModel "backend/internal/modules/payment-transaction/model"
	violationModel "backend/internal/modules/violation/model"
	violationrepo "backend/internal/modules/violation/repository"
	violationsvc "backend/internal/modules/violation/service"
	userdto "backend/internal/modules/user/dto"
	usermodel "backend/internal/modules/user/model"
	pagedto "backend/pkg/dto"

	"github.com/google/uuid"
)

type service struct {
	repo             repository.Repository
	violationRepo  violationrepo.Repository
	fineCalculator violationsvc.FineCalculationService
}

func NewService(repo repository.Repository, violationRepo violationrepo.Repository, fineCalculator violationsvc.FineCalculationService) Service {
	return &service{
		repo:           repo,
		violationRepo:  violationRepo,
		fineCalculator: fineCalculator,
	}
}

func (s *service) Create(req dto.CreateInvoiceRequest) error {
	violation, err := s.violationRepo.FindByID(req.ViolationID.String())
	if err != nil {
		return err
	}

	amount, err := s.fineCalculator.Calculate(&violation.FineRuleVersion)
	if err != nil {
		return err
	}

	invoice := invoiceModel.Invoice{
		BaseModel: dbmodel.BaseModel{
			ID: uuid.New(),
		},
		ViolationID: req.ViolationID,
		Amount:      amount,
		Status:      req.Status,
	}

	parsedMemberID, err := uuid.Parse(strings.TrimSpace(req.MemberID))
	if err != nil {
		return fmt.Errorf("invalid member_id: %w", err)
	}
	invoice.MemberID = parsedMemberID

	return s.repo.Create(&invoice)
}

func (s *service) GetByID(id string) (*dto.InvoiceResponse, error) {
	invoice, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return toResponse(invoice), nil
}

func (s *service) GetByMemberID(memberID string) ([]dto.InvoiceResponse, error) {
	invoices, err := s.repo.FindByMemberID(memberID)
	if err != nil {
		return nil, err
	}

	responses := make([]dto.InvoiceResponse, 0, len(invoices))
	for i := range invoices {
		responses = append(responses, *toResponse(&invoices[i]))
	}

	return responses, nil
}

func (s *service) GetAll(query pagedto.PaginationDTO, filter dto.ListInvoiceRequest) ([]dto.InvoiceResponse, int64, error) {
	query.Normalize()

	invoice, total, err := s.repo.FindAll(query, filter.Search, filter.Status)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.InvoiceResponse, 0, len(invoice))
	for i := range invoice {
		responses = append(responses, *toResponse(&invoice[i]))
	}

	return responses, total, nil
}

func (s *service) Update(id string, req dto.UpdateInvoiceRequest) error {
	invoice, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if req.ViolationID != nil && *req.ViolationID != uuid.Nil {
		invoice.ViolationID = *req.ViolationID
	}

	if req.MemberID != nil && strings.TrimSpace(*req.MemberID) != "" {
		parsedMemberID, err := uuid.Parse(strings.TrimSpace(*req.MemberID))
		if err != nil {
			return fmt.Errorf("invalid member_id: %w", err)
		}
		invoice.MemberID = parsedMemberID
	}

	if req.Status != nil && strings.TrimSpace(string(*req.Status)) != "" {
		invoice.Status = *req.Status
	}

	return s.repo.Update(invoice)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}

func toResponse(invoice *invoiceModel.Invoice) *dto.InvoiceResponse {
	return &dto.InvoiceResponse{
		ID:          invoice.ID.String(),
		ViolationID: invoice.ViolationID,
		Amount:      invoice.Amount,
		Status:      invoice.Status,
		Member:      toUserThrow2(invoice.Member),
		Violation:   toViolationThrow(invoice.Violation),
		Payment:     toPaymentThrow(invoice.Payment),
		CreatedAt:   invoice.CreatedAt,
		UpdatedAt:   invoice.UpdatedAt,
	}
}

func toViolationThrow(violation violationModel.Violation) dto.ViolationHistoryThrow {
	details := make([]dto.ViolationFineRuleDetailThrow, 0, len(violation.FineRuleVersion.Details))
	for i := range violation.FineRuleVersion.Details {
		detail := violation.FineRuleVersion.Details[i]
		details = append(details, dto.ViolationFineRuleDetailThrow{
			ID:            detail.ID.String(),
			RuleType:      detail.RuleType,
			Key:           detail.Key,
			Value:         detail.Value,
			RuleVersionID: detail.RuleVersionID.String(),
		})
	}

	return dto.ViolationHistoryThrow{
		ID:         violation.ID.String(),
		PlateNumber: violation.PlateNumber,
		Location:   violation.Location,
		OccurredAt: violation.OccurredAt,
		PhotoURL:   violation.PhotoURL,
		FineRuleVersion: dto.ViolationFineRuleVersionThrow{
			ID:            violation.FineRuleVersion.ID.String(),
			VersionNumber: violation.FineRuleVersion.VersionNumber,
			IsActive:      violation.FineRuleVersion.IsActive,
			Details:       details,
		},
	}
}

func toUserThrow2(user usermodel.User) userdto.UserThrow2 {
	return userdto.UserThrow2{
		ID:   user.ID.String(),
		Name: user.Name,
	}
}

func toPaymentThrow(payment paymentModel.PaymentTransaction) paymentdto.PaymentTransactionThrow {
	return paymentdto.PaymentTransactionThrow{
		ID:                    payment.ID.String(),
		InternalTransactionID: payment.InternalTransactionID,
		Amount:                payment.Amount,
		Status:                payment.Status,
		Scenario:              payment.Scenario,
		PaidAt:                payment.PaidAt,
	}
}
