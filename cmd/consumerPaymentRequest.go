package main

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	coreapi_midtrans "github.com/SyaibanAhmadRamadhan/go-midtrans-sdk/coreapi"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/payment-service/internal/conf"
	"github.com/mini-e-commerce-microservice/payment-service/internal/infra"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/outbox_events"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_methods"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_sources"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payments"
	"github.com/mini-e-commerce-microservice/payment-service/internal/services/order_saga"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

var consumerPaymentRequest = &cobra.Command{
	Use:   "consumerPaymentRequest",
	Short: "consumerPaymentRequest",
	Run: func(cmd *cobra.Command, args []string) {
		appConf := conf.LoadAppConf()
		kafkaConf := conf.LoadKafkaConf()
		otelConf := conf.LoadOtelConf()

		closeFnOtel := infra.NewOtel(otelConf, appConf.TracerName)
		kafkaBroker := ekafka.New(ekafka.WithOtel())
		pgdb, pgdbCloseFn := infra.NewPostgresql(appConf.DatabaseDsn)
		rdbms := wsqlx.NewRdbms(pgdb)
		midtransSdk := coreapi_midtrans.NewAPI(coreapi_midtrans.WithOtel(), coreapi_midtrans.ServerKey(appConf.PaymentGateway.Midtrans.ServerKey))

		outboxEventRepository := outbox_events.New(rdbms)
		paymentMethodRepository := payment_methods.New(rdbms)
		paymentSourceRepository := payment_sources.New(rdbms)
		paymentRepository := payments.New(rdbms)

		orderSagaService := order_saga.New(order_saga.Opt{
			KafkaConf:                kafkaConf,
			Kafka:                    kafkaBroker,
			PaymentRepository:        paymentRepository,
			PaymentSourceRepository:  paymentSourceRepository,
			PaymentMethodsRepository: paymentMethodRepository,
			OutboxEventsRepository:   outboxEventRepository,
			MidtransChargeEWallet:    midtransSdk,
			DatabaseTransaction:      rdbms,
		})

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		go func() {
			if err := orderSagaService.PaymentRequest(ctx); err != nil {
				log.Err(err).Msg("error payment request consumer")
				cancel()
			}
		}()

		<-ctx.Done()
		closeFnOtel(context.TODO())
		pgdbCloseFn(context.TODO())
		log.Info().Msg("Received shutdown signal, shutting down server gracefully...")
	},
}
