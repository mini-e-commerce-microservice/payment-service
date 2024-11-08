package primitive

type PaymentFeeType string

const (
	PaymentFeeTypePercentage PaymentFeeType = "percentage"
	PaymentFeeTypeFixed      PaymentFeeType = "fixed"
)
