package storage

import (
	"context"

	"github.com/uruddarraju/thyra/pkg/api/runtime"
)

// The limitation now is that there can only be one storage provider for all objects registered to the api gateway
// TODO: We would need to rethink this if we want to give users a way to register types with different backend storages
type ListOptions struct {
	APIGroup      string
	Type          string
	Name          string
	LabelSelector map[string]string
}

type Storage interface {
	RegisterGroup(ctx context.Context, group string) error
	UnregisterGroup(ctx context.Context, group string) error
	RegisterKind(ctx context.Context, group string, kind string) error
	UnregisterKind(ctx context.Context, group string, kind string) error
	List(ctx context.Context, options ListOptions) ([]api.Object, error)
	Get(ctx context.Context, lookup api.Object) (api.Object, error)
	Create(ctx context.Context, item api.Object) error
	Update(ctx context.Context, item api.Object, original api.Object) (api.Object, error)
	Delete(ctx context.Context, item api.Object) error
}
