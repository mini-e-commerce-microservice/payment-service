package models

type PaymentMethod struct {
	ID             int64   `db:"id"`
	IsActive       bool    `db:"is_active"`
	PaymentFee     float64 `db:"payment_fee"`
	PaymentFeeType string  `db:"payment_fee_type"`
	Code           string  `db:"code"`
	Category       string  `db:"category"`
	Name           string  `db:"name"`
	Image          string  `db:"image"`
}
