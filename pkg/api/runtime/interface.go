package runtime

import (
	"github.com/uruddarraju/thyra/pkg/api/types"
)

type Object interface {
	GetKind() string
	GetGroup() string
	GetMetadata() *api.ObjectMeta
}
