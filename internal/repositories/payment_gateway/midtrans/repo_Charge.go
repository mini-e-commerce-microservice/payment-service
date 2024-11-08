package midtrans

import (
	"context"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_gateway"
)

func (r *repository) Charge(ctx context.Context, input payment_gateway.ChargeInput) (output payment_gateway.ChargeOutput, err error) {
	return
}
