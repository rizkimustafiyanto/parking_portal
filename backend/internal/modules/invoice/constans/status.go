package constans

type InvoiceStatus string

const (
	InvoicePending  InvoiceStatus = "PENDING"
	InvoicePaid     InvoiceStatus = "PAID"
	InvoiceOverdue  InvoiceStatus = "OVERDUE"
	InvoiceCanceled InvoiceStatus = "CANCELED"
)