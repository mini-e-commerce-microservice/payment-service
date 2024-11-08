package payment_sources

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/payment-service/internal/models"
)

func (r *repository) FindOne(ctx context.Context, input FindOneInput) (output FindOneOutput, err error) {
	query := r.sq.Select("*").From("payment_sources").Limit(1)
	if input.Code.Valid {
		query = query.Where(squirrel.Eq{"code": input.Code.String})
	}

	err = r.rdbms.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeStruct, &output.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = collection.Err(err)
		}
		return output, err
	}

	return
}

type FindOneInput struct {
	Code null.String
}

type FindOneOutput struct {
	Data models.PaymentSource
}
