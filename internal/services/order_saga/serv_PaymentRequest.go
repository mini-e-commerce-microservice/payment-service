package order_saga

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"github.com/SyaibanAhmadRamadhan/go-midtrans-sdk"
	coreapi_midtrans "github.com/SyaibanAhmadRamadhan/go-midtrans-sdk/coreapi"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/payment-service/internal/models"
	"github.com/mini-e-commerce-microservice/payment-service/internal/primitive"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/outbox_events"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_methods"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payment_sources"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories/payments"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.22.0"
	"go.opentelemetry.io/otel/trace"
	"time"
)

func (s *service) PaymentRequest(ctx context.Context) (err error) {
	output, err := s.kafkaBroker.Subscribe(ctx, ekafka.SubInput{
		Config: kafka.ReaderConfig{
			Brokers: []string{s.kafkaConf.Host},
			GroupID: s.kafkaConf.Topic.OrderSaga.AggregatePaymentRequest.ConsumerGroup.Paymentsvc,
			Topic:   s.kafkaConf.Topic.OrderSaga.AggregatePaymentRequest.Name,
		},
	})
	if err != nil {
		return collection.Err(err)
	}

	for {
		data := PayloadCreateOrderProduct{}
		msg, err := output.Reader.FetchMessage(ctx, &data)
		if err != nil {
			return collection.Err(err)
		}

		carrier := ekafka.NewMsgCarrier(&msg)
		ctxConsumer := s.propagators.Extract(context.Background(), carrier)

		ctxConsumer, span := otel.Tracer("").Start(ctxConsumer, "process order saga, payment request.",
			trace.WithAttributes())

		outboxEventPayload := models.OutboxEvent{
			AggregateType: "payment",
			AggregateID:   "",
			Type:          "create-payment",
			TraceParent:   carrier.Get("traceparent"),
		}

		err = s.databaseTransaction.DoTxContext(ctxConsumer, &sql.TxOptions{Isolation: sql.LevelReadCommitted},
			func(ctx context.Context, tx wsqlx.Rdbms) (err error) {
				paymentSourceOutput, err := s.paymentSourceRepository.FindOne(ctx, payment_sources.FindOneInput{
					Code: null.StringFrom(string(primitive.PaymentSourceMidtrans)),
				})
				if err != nil {
					err = s.errorOrderSaga(ctx, outboxEventPayload, tx, data.OrderID, err, err.Error(), "failed get payment source")
					if err != nil {
						return collection.Err(err)
					}
				}

				paymentMethodOutput, err := s.paymentMethodsRepository.FindOne(ctx, payment_methods.FindOneInput{
					Code: null.StringFrom(data.PaymentMethodCode),
				})
				if err != nil {
					err = s.errorOrderSaga(ctx, outboxEventPayload, tx, data.OrderID, err, err.Error(), "failed get payment methods")
					if err != nil {
						return collection.Err(err)
					}
				}

				if paymentMethodOutput.Data.PaymentFeeType == string(primitive.PaymentFeeTypePercentage) {

				}

				chargeGopayOutput, err := s.midtransChargeEWallet.ChargeGoPay(ctx, coreapi_midtrans.ChargeGoPayInput{
					TransactionDetail: midtrans.TransactionDetail{
						OrderID:     fmt.Sprintf("%d", data.OrderID),
						GrossAmount: int64(data.TotalAmount),
					},
					CustomExpiry: &midtrans.CustomExpiry{
						ExpiryDuration: 2,
						Unit:           "minute",
					},
				})
				if err != nil || chargeGopayOutput.ErrorBadReqResponse != nil {
					msgError := ""
					if err != nil {
						msgError = err.Error()
					} else if chargeGopayOutput.ErrorBadReqResponse != nil {
						msgError = chargeGopayOutput.ErrorBadReqResponse.StatusMessage
					}

					err = s.errorOrderSaga(ctx, outboxEventPayload, tx, data.OrderID, err, msgError, "failed charge gopay out")
					if err != nil {
						return collection.Err(err)
					}
				}

				paymentAction := make([]models.PaymentAction, 0, len(chargeGopayOutput.ResponseSuccess.Actions))
				for _, action := range chargeGopayOutput.ResponseSuccess.Actions {
					paymentAction = append(paymentAction, models.PaymentAction{
						Name:   action.Name,
						URL:    action.URL,
						Method: action.Method,
					})
				}

				payment := models.Payment{
					PaymentSourceCode: paymentSourceOutput.Data.Code,
					PaymentMethodCode: paymentMethodOutput.Data.Code,
					Status:            string(primitive.PaymentStatusPending),
					FraudStatus:       "",
					ExpiredAt:         time.Now().UTC().Add(2 * time.Minute),
					TransactionAt:     null.Time{},
					OrderID:           data.OrderID,
					StatusMessage:     chargeGopayOutput.ResponseSuccess.StatusMessage,
					PaymentType:       chargeGopayOutput.ResponseSuccess.PaymentType,
					GrossAmount:       data.TotalAmount,
					SignatureKey:      chargeGopayOutput.ResponseSuccess.SignatureKey,
					Actions:           paymentAction,
					CreatedAt:         time.Now().UTC(),
					UpdatedAt:         time.Now().UTC(),
				}

				paymentCreateOutput, err := s.paymentRepository.Create(ctx, payments.CreateInput{
					Tx:   tx,
					Data: payment,
				})
				if err != nil {
					return collection.Err(err)
				}

				payment.ID = paymentCreateOutput.ID
				outboxEventPayload.Payload = models.PaymentPayloadOutboxEvent{
					PaymentData: payment,
				}
				outboxEventPayload.AggregateID = fmt.Sprintf("%d", payment.ID)
				_, err = s.outboxEventsRepository.Create(ctx, outbox_events.CreateInput{
					Tx:   tx,
					Data: outboxEventPayload,
				})
				if err != nil {
					return collection.Err(err)
				}

				return
			},
		)
		if err != nil {
			spanRecordErrorWIthEnd(span, err, "failed in transaction database")
			return collection.Err(err)
		}

		err = output.Reader.CommitMessages(ctx, msg)
		if err != nil {
			spanRecordErrorWIthEnd(span, err, "failed commit message")
			return collection.Err(err)
		}
		span.End()
	}
}

func spanRecordErrorWIthEnd(span trace.Span, err error, errType string) {
	span.RecordError(collection.Err(err))
	span.SetStatus(codes.Error, err.Error())
	span.SetAttributes(semconv.ErrorTypeKey.String(errType))
	span.End()
}

func (s *service) errorOrderSaga(ctx context.Context, outboxEvent models.OutboxEvent, tx wsqlx.Rdbms, orderID int64, errRecord error, errType, msgError string) (err error) {
	if !errors.Is(errRecord, repositories.ErrNoRecordFound) {
		return collection.Err(errRecord)
	}

	span := trace.SpanFromContext(ctx)
	span.RecordError(collection.Err(errRecord))
	span.SetStatus(codes.Error, msgError)
	span.SetAttributes(semconv.ErrorTypeKey.String(errType))

	outboxEvent.Payload = models.PaymentPayloadOutboxEvent{
		ErrorReason: &msgError,
		PaymentData: models.Payment{
			Status:  string(primitive.PaymentStatusFailed),
			OrderID: orderID,
		},
	}
	_, err = s.outboxEventsRepository.Create(ctx, outbox_events.CreateInput{
		Tx:   tx,
		Data: outboxEvent,
	})
	if err != nil {
		return collection.Err(err)
	}

	return
}
