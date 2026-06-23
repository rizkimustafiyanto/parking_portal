package service

import (
	"fmt"
	"strings"
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

	baseAmount, err := baseAmountForViolationType(strings.TrimSpace(violation.ViolationType))
	if err != nil {
		return 0, err
	}

	timeMultiplier, err := timeMultiplierForOccurredAt(violation.OccurredAt)
	if err != nil {
		return 0, err
	}

	repeatMultiplier := repeatMultiplierForCount(unpaidCount)

	return baseAmount * timeMultiplier * repeatMultiplier, nil
}

func baseAmountForViolationType(violationType string) (float64, error) {
	switch strings.ToLower(violationType) {
	case "expired_meter":
		return 50000, nil
	case "no_parking_zone":
		return 150000, nil
	case "blocking_hydrant":
		return 250000, nil
	case "disabled_spot":
		return 500000, nil
	default:
		return 0, fmt.Errorf("unsupported violation type: %s", violationType)
	}
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
