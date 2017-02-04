package local

import (
	"fmt"

	"github.com/uruddarraju/thyra/pkg/runtime"
	"golang.org/x/net/context"
	"github.com/uruddarraju/thyra/pkg/storage"
)

type LocalStorage struct {
	store map[string]map[string]runtime.Object
}

func (ls *LocalStorage) List(ctx context.Context, key string, listObj runtime.Object) error {
	objType := listObj.GetKind()
	typeFound, typeStore := ls.store[objType]
	if !typeFound {
		return storage.NewNotFoundError(fmt.Sprintf("Unable to find type %s", objType))
	}

	typeStore[key]
	return nil, nil
}

type Storage interface {
	Find(ctx context.Context, lookup runtime.Object) ([]runtime.Object, error)
	Create(ctx context.Context, item runtime.Object) error
	Update(ctx context.Context, item runtime.Object, original runtime.Object) error
	Delete(ctx context.Context, item runtime.Object) error
	Clear(ctx context.Context, lookup runtime.Object) (int, error)
}

