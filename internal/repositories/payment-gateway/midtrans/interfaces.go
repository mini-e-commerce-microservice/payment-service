package midtrans

import (
	"context"
	payment_gateway "github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment-gateway"
)

type Repository interface {
	Charge(ctx context.Context, input payment_gateway.ChargeInput) (output payment_gateway.ChargeOutput, err error)
}
