package service

import (
	"fmt"
	"time"

	"backend/internal/modules/violation/model"
)

type FineCalculationService interface {
	Calculate(version *model.FineRuleVersion, violation *model.Violation, unpaidCount int64) (float64, error)
}

type fineCalculationService struct{}

func NewFineCalculationService() FineCalculationService {
	return &fineCalculationService{}
}

func (s *fineCalculationService) Calculate(version *model.FineRuleVersion, violation *model.Violation, unpaidCount int64) (float64, error) {
	if version == nil {
		return 0, fmt.Errorf("fine rule version is required")
	}
	if violation == nil {
		return 0, fmt.Errorf("violation is required")
	}

	baseAmount := violation.ViolationType.BaseAmount
	if baseAmount <= 0 {
		return 0, fmt.Errorf("violation type %s has invalid base amount", violation.ViolationTypeCode)
	}

	timeMultiplier, err := timeMultiplierForOccurredAt(violation.OccurredAt)
	if err != nil {
		return 0, err
	}

	repeatMultiplier := repeatMultiplierForCount(unpaidCount)

	return baseAmount * timeMultiplier * repeatMultiplier, nil
}

func timeMultiplierForOccurredAt(occurredAt time.Time) (float64, error) {
	if occurredAt.IsZero() {
		return 0, fmt.Errorf("occurred_at is required")
	}

	hour := occurredAt.Hour()
	if hour >= 6 && hour < 22 {
		return 1.0, nil
	}

	return 1.5, nil
}

func repeatMultiplierForCount(unpaidCount int64) float64 {
	switch {
	case unpaidCount <= 0:
		return 1.0
	case unpaidCount == 1:
		return 1.5
	default:
		return 2.0
	}
}
