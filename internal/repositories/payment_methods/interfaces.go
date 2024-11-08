package payment_methods

import "context"

type Repository interface {
	FindOne(ctx context.Context, input FindOneInput) (output FindOneOutput, err error)
}
