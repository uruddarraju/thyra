package local

import (
	"errors"
	"fmt"

	"github.com/uruddarraju/thyra/pkg/runtime"
	"github.com/uruddarraju/thyra/pkg/storage"
	"golang.org/x/net/context"
)

type LocalStorage struct {
	// map of apiGroups, to types to Names of Objects to the Object itself
	// Example: map["foo.com"]["bar"]["namexyz"] = &XYZ{Name: "namexyz"} with *XYZ being of type runtime.Object
	store map[string]map[string]map[string]runtime.Object
}

func (ls *LocalStorage) RegisterGroup(ctx context.Context, group string) error {
	if ls.isGroupRegistered(group) {
		return nil
	}
	ls.store[group] = make(map[string]map[string]runtime.Object)
	return nil
}

func (ls *LocalStorage) UnregisterGroup(ctx context.Context, group string) error {
	groupStore, groupFound := ls.store[group]
	if groupFound {
		return nil
	}
	delete(groupStore, group)
	return nil
}

func (ls *LocalStorage) RegisterKind(ctx context.Context, group string, kind string) error {
	objGroup := group
	groupStore, groupFound := ls.store[objGroup]
	if !groupFound {
		return storage.NewNotFoundError(fmt.Sprintf("Unable to find group %s, check if it already registered to storage", objGroup))
	}

	objType := kind
	typeStore, typeFound := groupStore[objType]
	if typeFound {
		return nil
	}

	typeStore[kind] = make(map[string]runtime.Object)
	return nil
}

func (ls *LocalStorage) UnregisterKind(ctx context.Context, group string, kind string) error {
	objGroup := group
	groupStore, groupFound := ls.store[objGroup]
	if !groupFound {
		return storage.NewNotFoundError(fmt.Sprintf("Unable to find group %s, check if it already registered to storage", objGroup))
	}

	objType := kind
	typeStore, typeFound := groupStore[objType]
	if !typeFound {
		return nil
	}

	delete(typeStore, kind)
	return nil
}

func (ls *LocalStorage) List(ctx context.Context, options storage.ListOptions) ([]runtime.Object, error) {

	objGroup := options.APIGroup
	groupStore, groupFound := ls.store[objGroup]
	if !groupFound {
		return nil, storage.NewNotFoundError(fmt.Sprintf("Unable to find group %s, check if it already registered to storage", objGroup))
	}

	objType := options.Type
	typeStore, typeFound := groupStore[objType]
	if !typeFound {
		return nil, storage.NewNotFoundError(fmt.Sprintf("Unable to find type %s, check if it already registered to storage", objType))
	}

	// The storage should also make sure that all entries have unique name for a given object type
	// So, a name filter would trump any other filter provided in the options
	if len(options.Name) > 0 {
		result, exists := typeStore[options.Name]
		if !exists {
			return nil, storage.NewNotFoundError(fmt.Sprintf("Object %s/%s/%s not found", objGroup, objType, options.Name))
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

func (ls *LocalStorage) Get(ctx context.Context, lookup runtime.Object) (runtime.Object, error) {
	objGroup := lookup.GetGroup()
	groupStore, groupFound := ls.store[objGroup]
	if !groupFound {
		return nil, storage.NewNotFoundError(fmt.Sprintf("Unable to find group %s, check if it already registered to storage", objGroup))
	}

	objType := lookup.GetKind()
	typeStore, typeFound := groupStore[objType]
	if !typeFound {
		return nil, storage.NewNotFoundError(fmt.Sprintf("Unable to find type %s, check if it already registered to storage", objType))
	}

	// The storage should also make sure that all entries have unique name for a given object type
	// So, a name filter would trump any other filter provided in the options
	if len(lookup.GetName()) == 0 {
		return nil, errors.New("lookup object should have name to call the Get() function")
	}

	if result, exists := typeStore[lookup.GetName()]; !exists {
		return nil, storage.NewNotFoundError(fmt.Sprintf("Unable to find object with name %s", lookup.GetName()))
	} else {
		return result, nil
	}
}

func (ls *LocalStorage) Create(ctx context.Context, item runtime.Object) error {
	if err := ls.validateGroupKindRegistration(item.GetGroup(), item.GetKind()); err != nil {
		return err
	}

	if obj, _ := ls.Get(ctx, item); obj != nil {
		return storage.NewDuplicateEntryError(fmt.Sprintf("Object with name %s already exists", item.GetName()))
	}

	typeStore, _ := ls.store[item.GetGroup()][item.GetKind()]
	typeStore[item.GetName()] = item
	return nil
}

func (ls *LocalStorage) Update(ctx context.Context, item runtime.Object, original runtime.Object) (runtime.Object, error) {

	// TODO Should we check what attributes can be changed and not changed here or at the calling function ?
	if err := ls.validateGroupKindRegistration(original.GetGroup(), original.GetKind()); err != nil {
		return err
	}

	if obj, err := ls.Get(ctx, original); err != nil && storage.IsNotFoundError(err) {
		return storage.NewNotFoundError(fmt.Sprintf("Object with name %s does not exist", original.GetName()))
	} else if err != nil {
		return nil, err
	} else if obj == nil {
		return storage.NewNotFoundError(fmt.Sprintf("Object with name %s does not exist", original.GetName()))
	}

	typeStore, _ := ls.store[original.GetGroup()][original.GetKind()]
	typeStore[original.GetName()] = item
	return item, nil
}

func (ls *LocalStorage) Delete(ctx context.Context, item runtime.Object) error {
	if err := ls.validateGroupKindRegistration(item.GetGroup(), item.GetKind()); err != nil {
		return err
	}

	if obj, err := ls.Get(ctx, item); err != nil && storage.IsNotFoundError(err) {
		return storage.NewNotFoundError(fmt.Sprintf("Object with name %s does not exist", item.GetName()))
	} else if err != nil {
		return nil, err
	} else if obj == nil {
		return nil
	}

	typeStore, _ := ls.store[item.GetGroup()][item.GetKind()]
	typeStore[item.GetName()] = item
	return nil
}

// Each key provided in the filter should exist in the object metadata with values matching.
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

func (ls *LocalStorage) isGroupRegistered(group string) bool {
	_, groupFound := ls.store[group]
	if !groupFound {
		return false
	}
	return true
}

func (ls *LocalStorage) validateGroupKindRegistration(group, kind string) error {
	groupStore, groupFound := ls.store[group]
	if !groupFound {
		return storage.NewNotFoundError(fmt.Sprintf("Unable to find group %s, check if it already registered to storage", group))
	}

	_, typeFound := groupStore[kind]
	if !typeFound {
		return storage.NewNotFoundError(fmt.Sprintf("Unable to find type %s, check if it already registered to storage", kind))
	}

	return nil
}
