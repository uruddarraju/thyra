package storage

import (
	"context"

	"github.com/uruddarraju/thyra/pkg/runtime"
)

type Storage interface {
	Find(ctx context.Context, lookup runtime.Object) ([]runtime.Object, error)
	Create(ctx context.Context, item runtime.Object) error
	Update(ctx context.Context, item runtime.Object, original runtime.Object) error
	Delete(ctx context.Context, item runtime.Object) error
	Clear(ctx context.Context, lookup runtime.Object) (int, error)
}
