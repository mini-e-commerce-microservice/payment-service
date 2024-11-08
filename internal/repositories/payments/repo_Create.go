package payments

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

	actionMarshal, err := json.Marshal(input.Data.Actions)
	if err != nil {
		return output, collection.Err(err)
	}

	columns, values := collection.GetTagsWithValues(input.Data, "db", "id", "actions")
	columns = append(columns, "actions")
	values = append(values, actionMarshal)
	query := r.sq.Insert("payments").Columns(columns...).Values(values...).Suffix("RETURNING id")

	err = input.Tx.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeDefault, &output.ID)
	if err != nil {
		return output, collection.Err(err)
	}

	return
}

type CreateInput struct {
	Tx   wsqlx.ReadQuery
	Data models.Payment
}

type CreateOutput struct {
	ID int64
}
