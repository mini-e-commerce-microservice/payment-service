package order_saga

import "context"

type Service interface {
	PaymentRequest(ctx context.Context) (err error)
}
