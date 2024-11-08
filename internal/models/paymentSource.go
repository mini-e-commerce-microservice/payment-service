package models

type PaymentSource struct {
	ID          int    `db:"id"`
	IsActive    bool   `db:"is_active"`
	Code        string `db:"code"`
	Description string `db:"description"`
}
