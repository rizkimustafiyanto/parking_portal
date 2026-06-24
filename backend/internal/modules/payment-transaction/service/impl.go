package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"backend/internal/messaging"
	paymentconst "backend/internal/modules/payment-transaction/constans"
	"backend/internal/modules/payment-transaction/dto"
	paymentModel "backend/internal/modules/payment-transaction/model"
	"backend/internal/modules/payment-transaction/repository"
	pagedto "backend/pkg/dto"
)

type service struct {
	repo      repository.Repository
	publisher messaging.Publisher
}

func NewService(repo repository.Repository, publisher messaging.Publisher) Service {
	return &service{repo: repo, publisher: publisher}
}

func (s *service) Create(req dto.CreatePaymentTransactionRequest) error {
	if req.Status != paymentconst.PaymentSuccess || req.Scenario != paymentconst.ScenarioSuccess {
		return fmt.Errorf("member balance payment requires SUCCESS status and SUCCESS scenario")
	}

	if err := s.repo.ProcessMemberBalancePayment(
		req.InvoiceID.String(),
		req.InternalTransactionID,
		req.Amount,
		req.Status,
		req.Scenario,
		req.PaidAt,
	); err != nil {
		return err
	}

	if s.publisher != nil {
		if err := s.publisher.PublishJSON(context.Background(), "payment.completed", messaging.PaymentCompletedEvent{
			EventName:             "payment.completed",
			InvoiceID:             req.InvoiceID.String(),
			InternalTransactionID: req.InternalTransactionID,
			Amount:                req.Amount,
			Status:                string(req.Status),
			Scenario:              string(req.Scenario),
			PaidAt:                req.PaidAt,
			CreatedAt:             req.PaidAt,
		}); err != nil {
			log.Printf("publish payment.completed failed: %v", err)
		}
	}

	return nil
}

func (s *service) GetByID(id string) (*dto.PaymentTransactionResponse, error) {
	payments, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return toResponse(payments), nil
}

func (s *service) GetAll(query pagedto.PaginationDTO, filter dto.ListPaymentTransactionRequest) ([]dto.PaymentTransactionResponse, int64, error) {
	query.Normalize()

	payments, total, err := s.repo.FindAll(query, filter.Search, filter.Status, filter.Scenario)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]dto.PaymentTransactionResponse, 0, len(payments))
	for i := range payments {
		responses = append(responses, *toResponse(&payments[i]))
	}

	return responses, total, nil
}

func (s *service) Update(id string, req dto.UpdatePaymentTransactionRequest) error {
	payments, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if req.InternalTransactionID != nil && strings.TrimSpace(*req.InternalTransactionID) != "" {
		payments.InternalTransactionID = *req.InternalTransactionID
	}

	if req.Amount != nil {
		payments.Amount = *req.Amount
	}

	if req.Status != nil && strings.TrimSpace(string(*req.Status)) != "" {
		payments.Status = *req.Status
	}

	if req.Scenario != nil && strings.TrimSpace(string(*req.Scenario)) != "" {
		payments.Scenario = *req.Scenario
	}

	if req.PaidAt != nil && !req.PaidAt.IsZero() {
		payments.PaidAt = *req.PaidAt
	}

	return s.repo.Update(payments)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}

func toResponse(payments *paymentModel.PaymentTransaction) *dto.PaymentTransactionResponse {
	return &dto.PaymentTransactionResponse{
		ID:                    payments.ID.String(),
		InvoiceID:             payments.InvoiceID,
		InternalTransactionID: payments.InternalTransactionID,
		Amount:                payments.Amount,
		Status:                payments.Status,
		Scenario:              payments.Scenario,
		CreatedAt:             payments.CreatedAt,
	}
}
