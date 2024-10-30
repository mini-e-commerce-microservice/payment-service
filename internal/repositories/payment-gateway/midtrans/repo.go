package midtrans

import "github.com/mini-e-commerce-microservice/payment-service/generated/proto/secret_proto"

type repository struct {
	conf *secret_proto.PaymentServicePaymentGatewayMidtrans
}

func New(conf *secret_proto.PaymentServicePaymentGatewayMidtrans) *repository {
	return &repository{
		conf: conf,
	}
}
