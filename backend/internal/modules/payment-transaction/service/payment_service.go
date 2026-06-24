package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type mockPaymentService struct{}

func NewPaymentService() PaymentService {
	return &mockPaymentService{}
}

func (s *mockPaymentService) Charge(invoiceID string, amount float64, scenario string) ChargeResult {
	scenario = strings.ToUpper(strings.TrimSpace(scenario))

	status := "failed"
	if scenario == "SUCCESS" {
		status = "paid"
	}

	return ChargeResult{
		Status:        status,
		TransactionID: buildTransactionID(invoiceID, amount, status),
	}
}

func buildTransactionID(invoiceID string, amount float64, status string) string {
	token := make([]byte, 4)
	if _, err := rand.Read(token); err != nil {
		token = []byte(time.Now().Format("150405"))
	}

	safeInvoiceID := strings.ReplaceAll(strings.ReplaceAll(invoiceID, "-", ""), " ", "")
	if len(safeInvoiceID) > 8 {
		safeInvoiceID = safeInvoiceID[:8]
	}

	return fmt.Sprintf("TRX-%s-%s-%.0f-%s", strings.ToUpper(status), safeInvoiceID, amount, hex.EncodeToString(token))
}
