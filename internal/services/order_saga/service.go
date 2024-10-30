package order_saga

import (
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	coreapi_midtrans "github.com/SyaibanAhmadRamadhan/go-midtrans-sdk/coreapi"
	"github.com/mini-e-commerce-microservice/payment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/outbox_events"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_methods"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_sources"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payments"
)

type service struct {
	conf                    *secret_proto.Kafka
	kafka                   ekafka.KafkaPubSub
	paymentGateway          coreapi_midtrans.ChargeEWalletAPI
	paymentRepository       payments.Repository
	paymentSourceRepository payment_sources.Repository
	paymentMethods          payment_methods.Repository
	outboxEventsRepository  outbox_events.Repository
}

type Opt struct {
	KafkaConf               *secret_proto.Kafka
	Kafka                   ekafka.KafkaPubSub
	PaymentGateway          coreapi_midtrans.ChargeEWalletAPI
	PaymentRepository       payments.Repository
	PaymentSourceRepository payment_sources.Repository
	PaymentMethods          payment_methods.Repository
	OutboxEventsRepository  outbox_events.Repository
}

func New(opt Opt) *service {
	return &service{
		conf:                    opt.KafkaConf,
		kafka:                   opt.Kafka,
		paymentGateway:          opt.PaymentGateway,
		paymentRepository:       opt.PaymentRepository,
		paymentSourceRepository: opt.PaymentSourceRepository,
		paymentMethods:          opt.PaymentMethods,
		outboxEventsRepository:  opt.OutboxEventsRepository,
	}
}
