package repo

import "context"

type Auth interface {
	Create(ctx context.Context) error
}
