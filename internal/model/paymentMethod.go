package model

type PaymentMethod struct {
	ID       int64  `db:"id"`
	IsActive bool   `db:"is_active"`
	Code     string `db:"code"`
	Category string `db:"category"`
	Name     string `db:"name"`
	Image    string `db:"image"`
}
