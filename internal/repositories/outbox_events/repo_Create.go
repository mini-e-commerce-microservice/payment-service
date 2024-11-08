package outbox_events

import (
	"context"
	"encoding/json"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/payment-service/internal/models"
	"github.com/mini-e-commerce-microservice/payment-service/internal/repositories"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	if input.Tx == nil {
		return output, collection.Err(repositories.ErrTxIsNil)
	}

	payloadMarshal, err := json.Marshal(input.Data.Payload)
	if err != nil {
		return output, collection.Err(err)
	}

	query := r.sq.Insert("outbox_events").Columns("aggregatetype", "aggregateid", "type", "payload", "trace_parent").
		Values(input.Data.AggregateType, input.Data.AggregateID, input.Data.Type, string(payloadMarshal), input.Data.TraceParent).Suffix("RETURNING id")

	err = input.Tx.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeDefault, &output.ID)
	if err != nil {
		return output, collection.Err(err)
	}

	return
}

type CreateInput struct {
	Tx   wsqlx.ReadQuery
	Data models.OutboxEvent
}

type CreateOutput struct {
	ID string
}
