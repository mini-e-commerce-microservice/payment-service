package order_saga

import (
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	coreapi_midtrans "github.com/SyaibanAhmadRamadhan/go-midtrans-sdk/coreapi"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/payment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/outbox_events"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_methods"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_sources"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payments"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type service struct {
	kafkaConf                *secret_proto.Kafka
	kafkaBroker              ekafka.KafkaPubSub
	paymentRepository        payments.Repository
	paymentSourceRepository  payment_sources.Repository
	paymentMethodsRepository payment_methods.Repository
	outboxEventsRepository   outbox_events.Repository
	propagators              propagation.TextMapPropagator
	midtransChargeEWallet    coreapi_midtrans.ChargeEWalletAPI
	databaseTransaction      wsqlx.Tx
}

type Opt struct {
	KafkaConf                *secret_proto.Kafka
	Kafka                    ekafka.KafkaPubSub
	PaymentRepository        payments.Repository
	PaymentSourceRepository  payment_sources.Repository
	PaymentMethodsRepository payment_methods.Repository
	OutboxEventsRepository   outbox_events.Repository
	MidtransChargeEWallet    coreapi_midtrans.ChargeEWalletAPI
	DatabaseTransaction      wsqlx.Tx
}

func New(opt Opt) *service {
	return &service{
		propagators:              otel.GetTextMapPropagator(),
		kafkaConf:                opt.KafkaConf,
		kafkaBroker:              opt.Kafka,
		paymentRepository:        opt.PaymentRepository,
		paymentSourceRepository:  opt.PaymentSourceRepository,
		paymentMethodsRepository: opt.PaymentMethodsRepository,
		outboxEventsRepository:   opt.OutboxEventsRepository,
		databaseTransaction:      opt.DatabaseTransaction,
		midtransChargeEWallet:    opt.MidtransChargeEWallet,
	}
}
