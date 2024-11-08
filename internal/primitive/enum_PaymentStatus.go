package primitive

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusRejected PaymentStatus = "REJECTED"
	PaymentStatusSettle   PaymentStatus = "SETTLED"
	PaymentStatusFailed   PaymentStatus = "FAILED"
)
