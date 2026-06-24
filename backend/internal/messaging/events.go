package messaging

import "time"

type InvoiceCreatedEvent struct {
	EventName   string    `json:"event_name"`
	InvoiceID   string    `json:"invoice_id"`
	MemberID    string    `json:"member_id"`
	ViolationID string    `json:"violation_id"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaymentCompletedEvent struct {
	EventName             string    `json:"event_name"`
	PaymentID             string    `json:"payment_id"`
	InvoiceID             string    `json:"invoice_id"`
	InternalTransactionID string    `json:"internal_transaction_id"`
	Amount                float64   `json:"amount"`
	Status                string    `json:"status"`
	Scenario              string    `json:"scenario"`
	PaidAt                time.Time `json:"paid_at"`
	CreatedAt             time.Time `json:"created_at"`
}

