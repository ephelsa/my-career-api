package documenttype

import "context"

type Repository interface {
	FetchAll(ctx context.Context) ([]DocumentType, error)
	New(ctx context.Context, dt *DocumentType) error
	//DeleteByID(ctx context.Context, id string) error
}
