package core_transaction

import "context"

type Manager interface {
	WithTransaction(
		ctx context.Context,
		fn func(ctx context.Context) error,
	) error
}
