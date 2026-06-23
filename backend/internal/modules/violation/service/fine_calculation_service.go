package service

import (
	"fmt"
	"strconv"
	"strings"

	"backend/internal/modules/violation/model"
)

type FineCalculationService interface {
	Calculate(version *model.FineRuleVersion) (float64, error)
}

type fineCalculationService struct{}

func NewFineCalculationService() FineCalculationService {
	return &fineCalculationService{}
}

func (s *fineCalculationService) Calculate(version *model.FineRuleVersion) (float64, error) {
	if version == nil {
		return 0, fmt.Errorf("fine rule version is required")
	}

	var total float64
	var hasAmount bool

	for i := range version.Details {
		detail := version.Details[i]
		key := strings.ToLower(strings.TrimSpace(detail.Key))
		value := strings.TrimSpace(detail.Value)

		numericValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
		}

		switch key {
		case "amount", "fine_amount", "base_amount":
			total += numericValue
			hasAmount = true
		case "percentage", "rate":
			total += numericValue
			hasAmount = true
		default:
			total += numericValue
			hasAmount = true
		}
	}

	if !hasAmount {
		return 0, fmt.Errorf("fine rule version %s has no numeric details", version.ID.String())
	}

	return total, nil
}
