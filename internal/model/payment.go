package model

import "time"

type Payment struct {
	ID                int             `db:"id"`
	PaymentSourceCode string          `db:"payment_source_code"`
	PaymentMethodCode string          `db:"payment_method_code"`
	Status            string          `db:"status"`
	FraudStatus       string          `db:"fraud_status"`
	ExpiredAt         time.Time       `db:"expired_at"`
	OrderID           string          `db:"order_id"`
	StatusMessage     string          `db:"status_message"`
	PaymentType       string          `db:"payment_type"`
	GrossAmount       float64         `db:"gross_amount"`
	SignatureKey      string          `db:"signature_key"`
	Actions           []PaymentAction `db:"actions"`
	CreatedAt         time.Time       `db:"created_at"`
	UpdatedAt         time.Time       `db:"updated_at"`
}

type PaymentAction struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Method string `json:"method"`
}
