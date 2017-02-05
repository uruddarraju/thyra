package storage

import (
	"context"

	"github.com/uruddarraju/thyra/pkg/runtime"
)

// The limitation now is that there can only be one storage provider for all objects registered to the api gateway
// TODO: We would need to rethink this if we want to give users a way to register types with different backend storages
type ListOptions struct {
	Name          string
	Type          string
	LabelSelector map[string]string
}

type Storage interface {
	List(ctx context.Context, options ListOptions) ([]runtime.Object, error)
	Get(ctx context.Context, lookup runtime.Object) (runtime.Object, error)
	Create(ctx context.Context, item runtime.Object) error
	Update(ctx context.Context, item runtime.Object, original runtime.Object) (runtime.Object, error)
	Delete(ctx context.Context, item runtime.Object) error
}
