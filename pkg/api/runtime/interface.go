package runtime

import (
	"github.com/uruddarraju/thyra/pkg/apis/thyra"
)

type Object interface {
	GetKind() string
	GetGroup() string
	GetMetadata() *thyra.ObjectMeta
}
