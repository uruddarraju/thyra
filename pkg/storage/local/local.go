package local

import (
	"fmt"

	"github.com/uruddarraju/thyra/pkg/runtime"
	"github.com/uruddarraju/thyra/pkg/storage"
	"golang.org/x/net/context"
)

type LocalStorage struct {
	store map[string]map[string]runtime.Object
}

func (ls *LocalStorage) List(ctx context.Context, options ListOptions) ([]runtime.Object, error) {
	objType := options.Type
	typeStore, typeFound := ls.store[objType]
	if !typeFound {
		return nil, storage.NewNotFoundError(fmt.Sprintf("Unable to find type %s, check if it already registered to storage", objType))
	}

	// The storage should also make sure that all entries have unique name for a given object type
	// So, a name filter would trump any other filter provided in the options
	if len(options.Name) > 0 {
		result, exists := typeStore[options.Name]
		if !exists {
			return nil, storage.NewNotFoundError(fmt.Sprintf("Object not found"))
		}
		return result, nil
	}

	// Very inefficient to do Lists as we are iterating on maps, but for local, this should be ok.
	result := make([]runtime.Object, 0)
	for _, v := range typeStore {
		if ls.filter(v, options.LabelSelector) {
			append(result, v)
		}
	}

	return result
}

// Each keyname provided in the filter should exist in the object metadata with values matching.
// Object metadata can have more key-value pairs than the selector.
func (ls *LocalStorage) filter(object runtime.Object, selector map[string]string) bool {
	for key, value := range selector {
		if object.GetMetadata() == nil {
			return false
		}
		if objectValue, exists := object.GetMetadata()[key]; exists && value == objectValue {
			continue
		} else {
			return false
		}
	}
	return true
}

// Defined in interface, remove this after you finished implementing the local storage
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

// Until here
